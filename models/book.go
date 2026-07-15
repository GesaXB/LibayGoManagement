package models

type Book struct {
	BaseModel
	Title       string `gorm:"not null;size:255"`
	Isbnd       string `gorm:"not null"`
	Description string `gorm:"type:text"`
	Stock       uint   `gorm:"default:0"`

	AuthorId uint `json:"author_id"`
	Author   Author

	CategoryId uint `json:"category_id"`
	Category   Category

	Borrows []Borrow `gorm:"foreignKey:BookId"`
}
