package comment

type Comment struct {
	ID        uint   `json:"id"`
	Text      string `json:"text"`
	Date      int64  `json:"date"`
	UserEmail string `json:"user_email"`
	PostID    uint
}
