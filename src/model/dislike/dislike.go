package dislike

type Dislike struct {
	ID        uint   `json:"id"`
	UserEmail string `json:"user_email"`
	PostID    uint
}
