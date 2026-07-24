package models

import "time"

type Borrow struct {
	BaseModel

	UserId string `json:"user_id"`
	User   User

	BookId string `json:"book_id"`
	Book   Book

	BorrowDate time.Time  `json:"borrow_date"`
	ReturnDate *time.Time `json:"return_date"`

	Status string `gorm:"default:borrowed" json:"status"`
}
