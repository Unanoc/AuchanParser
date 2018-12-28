package router

import (
	"server/handlers"

	"github.com/buaazp/fasthttprouter"
)

func NewRouter() *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.GET("/api/products/product/:product_id/", handlers.GetProductByIdHandler) // получение определенного продукта по product_id
	router.GET("/api/products/all", handlers.GetProductsAllHandler) // получение всех продуктов
	router.GET("/api/products/status", handlers.GetProductsStatusHandler) // получение всех продуктов
	router.POST("/api/products/product/", handlers.PostProductByIdHandler) // создание продукта


	return router
}