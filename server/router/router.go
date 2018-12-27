package router

import (
	"server/handlers"

	"github.com/buaazp/fasthttprouter"
)

func NewRouter() *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.GET("/api/products/:product_id/", handlers.GetProductByIdHandler) // получение продукта по product_id
	// router.GET("/api/forum/:slug/threads", forum.ForumGetThreadsHandler) //done
	// router.GET("/api/forum/:slug/users", forum.ForumGetUsersHandler) //done

	return router
}