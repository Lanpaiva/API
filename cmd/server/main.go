package main

import "github.com/lanpaiva/api/configs"

func main() {
	configs, _ := configs.LoadConfig(".")
	println(configs.DBDriver)
}
