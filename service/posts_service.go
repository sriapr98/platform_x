package service

import (
	"platform_x/models/db"
	"platform_x/models/request"
	"platform_x/models/response"
	"platform_x/repository"
	"time"
)

type PostsService interface {
	Create(post request.CreatePost)
	GetAll() []response.Post

	GetById(postId int) (db.Post, error)
}

type postsService struct {
	postsRepository repository.PostsRepository
	likesRepository repository.LikesRepository
}

func (p postsService) Create(post request.CreatePost) {
	p.postsRepository.Create(db.Post{
		Id:        post.Id,
		Content:   post.Content,
		CreatedAt: time.Now(),
		CreatedBy: post.UserName,
	})
}

func (p postsService) GetById(postId int) (db.Post, error) {
	return p.postsRepository.GetById(postId)
}

func (p postsService) GetAll() []response.Post {
	var postsResponse []response.Post

	posts := p.postsRepository.GetAll()

	for _, post := range posts {
		likes := p.likesRepository.GetCount(post.Id)
		postsResponse = append(postsResponse, getPost(post, likes))
	}

	return postsResponse
}

func getPost(post db.Post, likes int) response.Post {
	return response.Post{
		Id:         post.Id,
		Content:    post.Content,
		CreatedAt:  post.CreatedAt,
		LikesCount: likes,
	}
}

func NewPostsService(postsRepository repository.PostsRepository, likesRepository repository.LikesRepository) PostsService {
	return postsService{postsRepository, likesRepository}
}
