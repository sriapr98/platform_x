package db

import "time"

type Post struct {
	Id        int
	Content   string
	CreatedAt time.Time
	CreatedBy string
}
