package repository

import (
	"platform_x/models/db"
)

type LikesRepository interface {
	Add(postId int, like db.Like)
	GetCount(postId int) int
}

type likesRepository struct {
	likes map[int][]db.Like
}

func (l *likesRepository) Add(postId int, like db.Like) {
	l.likes[postId] = append(l.likes[postId], like)
}

func (l *likesRepository) GetCount(postId int) int {
	return len(l.likes[postId])
}

func NewLikesRepository(likes map[int][]db.Like) LikesRepository {
	return &likesRepository{likes}
}
