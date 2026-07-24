package responsedto

type BookResponse struct {
	ID          uint            `json:"id"`
	Title       string          `json:"title"`
	Isbnd       string          `json:"isbnd"`
	Description string          `json:"description"`
	Stock       uint            `json:"stock"`
	Author      AuthorResponse  `json:"author"`
	Category    CategoryRespone `json:"category"`
}
