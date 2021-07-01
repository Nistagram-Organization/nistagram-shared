package like

type Like struct {
	ID        uint   `json:"id"`
	UserEmail string `json:"user_email"`
	PostID    uint
}
