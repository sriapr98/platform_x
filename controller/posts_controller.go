package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"platform_x/models/request"
	"platform_x/service"
)

type PostsController interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type postsController struct {
	postsService service.PostsService
}

func (l postsController) Create(ctx *gin.Context) {
	var createPostRequest request.CreatePost

	err := ctx.ShouldBindBodyWith(&createPostRequest, binding.JSON)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error_code": "ERR_INVALID_REQUEST", "error_message": "Invalid request"})
		return
	}

	l.postsService.Create(createPostRequest)

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"status": "Success"})
}

func (l postsController) Get(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, l.postsService.GetAll())
}

func NewPostsController(postsService service.PostsService) PostsController {
	return postsController{postsService}
}
