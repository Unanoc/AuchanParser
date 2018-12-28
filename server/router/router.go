package router

import (
	"server/handlers"

	"github.com/buaazp/fasthttprouter"
)

func NewRouter() *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.GET("/api/products/product/:product_id/", handlers.GetProductByIdHandler) // получение определенного продукта по product_id 
	router.GET("/api/products/all", handlers.GetProductsAllHandler) // получение всех продуктов (http://localhost:3000/api/products/all?limit=10&from=100&to=1000)
	router.GET("/api/products/status", handlers.GetProductsStatusHandler) // получение информации о количество продуктов в бд
	router.POST("/api/products/product/", handlers.PostProductByIdHandler) // создание продукта

	return router
}