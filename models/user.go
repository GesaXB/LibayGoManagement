package models

const (
	RoleAdmin = "Admin"
	RoleUser  = "User"
)

type User struct {
	BaseModel
	Name     string `gorm:"not null;size:100"`
	Email    string `gorm:"uniqueIndex;size:255;not null"`
	Password string `gorm:"size:255;not null"`
	Role     string `gorm:"default:user"`

	Borrows []Borrow `gorm:"foreingKey:UserId"`
}
