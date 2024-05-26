package Models

type Comment struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Text   string `json:"text"`
	PageID int    `json:"pageId"`
}
