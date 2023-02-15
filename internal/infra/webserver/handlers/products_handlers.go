package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lanpaiva/api/internal/dto"
	"github.com/lanpaiva/api/internal/entity"
	"github.com/lanpaiva/api/internal/infra/database"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

// banco de dados
func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Cria o produto
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	//recebe produto puro do arquivo DTO
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//CRIA DE FATO O PRODUTO
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Cria o produto e vincula no banco de dados
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}
