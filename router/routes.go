package router

import (
	"github.com/gin-gonic/gin"
	"platform_x/controller"
	"platform_x/models/db"
	"platform_x/repository"
	"platform_x/service"
)

func InitObjects() *gin.Engine {
	router := gin.Default()

	likesRepository := repository.NewLikesRepository(make(map[int][]db.Like))
	postsRepository := repository.NewPostsRepository(nil)

	postsService := service.NewPostsService(postsRepository, likesRepository)
	likesService := service.NewLikesService(likesRepository, postsService)

	likesController := controller.NewLikesController(likesService)
	postsController := controller.NewPostsController(postsService)

	postsEndpoints := router.Group("/posts")
	{
		postsEndpoints.POST("", postsController.Create)
		postsEndpoints.GET("", postsController.Get)
	}

	likesEndpoints := router.Group("/like")
	{
		likesEndpoints.POST("", likesController.Add)
	}

	return router
}
