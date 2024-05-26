package Models

type Moderation struct {
	ID        int  `json:"id" gorm:"primaryKey"`
	PageID    int  `json:"pageId"`
	CommentID int  `json:"commentId"`
	Flagged   bool `json:"approved"`
}
