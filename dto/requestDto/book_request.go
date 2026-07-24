package requestdto

type BookRequest struct {
	Title       string `json:"title" binding:"required"`
	Isbnd       string `json:"isbnd" binding:"required"`
	Description string `json:"description"`
	Stock       uint   `json:"stock"`
	AuthorId    uint   `json:"author_id" binding:"required"`
	CategoryId  uint   `json:"category_id" binding:"required"`
}
