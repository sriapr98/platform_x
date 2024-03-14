package response

import "time"

type Posts struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	Id         int       `json:"id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	LikesCount int       `json:"likes_count"`
}
