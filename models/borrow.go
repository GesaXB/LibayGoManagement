package models

import "time"

type Borrow struct {
	BaseModel

	UserId uint `json:"user_id"`
	User   User

	BookId uint `json:"book_id"`
	Book   Book

	BorrowDate time.Time  `json:"borrow_date"`
	ReturnDate *time.Time `json:"return_date"`

	Status string `gorm:"default:borrowed" json:"status"`
}
