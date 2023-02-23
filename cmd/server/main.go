package main

import (
	"log"
	"net/http"

	"github.com/lanpaiva/api/configs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/lanpaiva/api/internal/entity"
	"github.com/lanpaiva/api/internal/infra/database"
	"github.com/lanpaiva/api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHand := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHand := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// r.Use(LogRequest)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExperesIn", configs.JwtExperesIn))
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHand.CreateProduct)
		r.Get("/", productHand.FindAllProducts)
		r.Get("/{id}", productHand.GetProduct)
		r.Put("/{id}", productHand.UpdateProduct)
		r.Delete("/{id}", productHand.DeleteProduct)
	})

	r.Post("/users", userHand.CreateUser)
	r.Post("/users/generate_token", userHand.GetJWT)

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request, %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
