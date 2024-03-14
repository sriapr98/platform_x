package service

import (
	"platform_x/models/db"
	"platform_x/models/request"
	"platform_x/repository"
	"time"
)

type LikesService interface {
	Add(like request.Like) error
}

type likesService struct {
	likesRepository repository.LikesRepository
	postsService    PostsService
}

func (l likesService) Add(like request.Like) error {
	_, err := l.postsService.GetById(like.PostId)
	if err != nil {
		return err
	}

	l.likesRepository.Add(like.PostId, getLike(like))

	return nil
}

func getLike(like request.Like) db.Like {
	return db.Like{
		CreatedBy: like.UserName,
		CreatedAt: time.Now(),
	}
}

func NewLikesService(likesRepository repository.LikesRepository, postsService PostsService) LikesService {
	return likesService{likesRepository, postsService}
}
