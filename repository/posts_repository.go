package repository

import (
	"errors"
	"platform_x/models/db"
)

type PostsRepository interface {
	Create(post db.Post)
	GetAll() []db.Post
	GetById(postId int) (db.Post, error)
}

type postsRepository struct {
	posts []db.Post
}

func (p *postsRepository) GetAll() []db.Post {
	return p.posts
}

func (p *postsRepository) GetById(postId int) (db.Post, error) {
	for _, post := range p.posts {
		if post.Id == postId {
			return post, nil
		}
	}
	return db.Post{}, errors.New("no posts found")
}

func (p *postsRepository) Create(post db.Post) {
	p.posts = append(p.posts, post)
}

func NewPostsRepository(posts []db.Post) PostsRepository {
	return &postsRepository{posts}
}
