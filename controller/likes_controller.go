package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"platform_x/models/request"
	"platform_x/service"
)

type LikesController interface {
	Add(ctx *gin.Context)
}

type likesController struct {
	likesService service.LikesService
}

func (l likesController) Add(ctx *gin.Context) {
	var likeRequest request.Like

	err := ctx.ShouldBindBodyWith(&likeRequest, binding.JSON)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error_code": "ERR_INVALID_REQUEST", "error_message": "Invalid request"})
		return
	}

	err = l.likesService.Add(likeRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error_code": "ERR_INVALID_REQUEST", "error_message": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"status": "Success"})
}

func NewLikesController(likesService service.LikesService) LikesController {
	return likesController{likesService}
}
