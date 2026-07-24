package requestdto

type BookRequest struct {
	Title       string `json:"title" binding:"required"`
	Isbnd       string `json:"isbnd" binding:"required"`
	Description string `json:"description"`
	Stock       uint   `json:"stock"`
	AuthorId    string `json:"author_id" binding:"required"`
	CategoryId  string `json:"category_id" binding:"required"`
}
