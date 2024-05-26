package Models

type Recommendation struct {
	ID       string `json:"page_id"`
	Distance string `json:"vector_distance"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}
