package router

import (
	"server/handlers"

	"github.com/buaazp/fasthttprouter"
)

func NewRouter() *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.GET("/api/products/product/:product_id/", handlers.GetProductByIdHandler) // получение определенного продукта по product_id
	router.POST("/api/products/product/", handlers.PostProductByIdHandler) // создание продукта
	router.GET("/api/products/all", handlers.GetProductsAllHandler) // получение всех продуктов


	return router
}