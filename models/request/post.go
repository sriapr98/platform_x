package request

type CreatePost struct {
	Id       int    `json:"id"`
	Content  string `json:"content"`
	UserName string `json:"user_name"`
}

type FetchPosts struct {
	UserName string `json:"user_name"`
}
