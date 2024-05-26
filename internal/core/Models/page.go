package Models

type Page struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
