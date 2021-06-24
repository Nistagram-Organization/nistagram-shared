package dislike

type Dislike struct {
	ID     uint `json:"id"`
	PostID uint
	UserID uint
}
