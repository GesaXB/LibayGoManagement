package models

type Author struct {
	BaseModel
	Name string `gorm:"not null;size:255"`
	Bio  string `gorm:"size:255"`

	Book []Book `gorm:"foreignKey:AuthorId"`
}
