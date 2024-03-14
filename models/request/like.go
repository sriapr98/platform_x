package request

type Like struct {
	PostId   int    `json:"post_id"`
	UserName string `json:"user_name"`
}
