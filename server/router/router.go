package router

import (

	"github.com/buaazp/fasthttprouter"
)

func NewRouter() *fasthttprouter.Router {
	router := fasthttprouter.New()

	// router.POST("/api/forum/:slug", forum.ForumCreateHandler) //done
	// router.POST("/api/forum/:slug/create", forum.ForumCreateThreadHandler) //done
	// router.GET("/api/forum/:slug/details", forum.ForumGetOneHandler) //done
	// router.GET("/api/forum/:slug/threads", forum.ForumGetThreadsHandler) //done
	// router.GET("/api/forum/:slug/users", forum.ForumGetUsersHandler) //done

	// router.GET("/api/post/:id/details", post.PostGetOneHandler) //done
	// router.POST("/api/post/:id/details", post.PostUpdateHandler) //done

	// router.GET("/api/service/status", service.StatusHandler) //done
	// router.POST("/api/service/clear", service.ClearHandler) //done

	// router.GET("/api/thread/:slug_or_id/details", thread.ThreadGetOneHandler) //done
	// router.GET("/api/thread/:slug_or_id/posts", thread.ThreadGetPostsHandler)
	// router.POST("/api/thread/:slug_or_id/create", thread.ThreadCreateHandler) //done
	// router.POST("/api/thread/:slug_or_id/details", thread.ThreadUpdateHandler) // done
	// router.POST("/api/thread/:slug_or_id/vote", thread.ThreadVoteHandler) //done

	// router.GET("/api/user/:nickname/profile", user.UserGetOneHandler) //done
	// router.POST("/api/user/:nickname/create", user.UserCreateHandler) //done
	// router.POST("/api/user/:nickname/profile", user.UserUpdateHandler) //done

	return router
}