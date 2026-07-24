package models

type Book struct {
	BaseModel
	Title       string `gorm:"not null;size:255;unique"`
	Isbnd       string `gorm:"not null"`
	Description string `gorm:"type:text"`
	Stock       uint   `gorm:"default:0"`

	AuthorId string `json:"author_id"`
	Author   Author

	CategoryId string `json:"category_id"`
	Category   Category

	Borrows []Borrow `gorm:"foreignKey:BookId"`
}
