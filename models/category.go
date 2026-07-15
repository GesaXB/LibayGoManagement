package models

type Category struct {
	BaseModel
	Name string `gorm:"not null;size:255"`

	Book []Book `gorm:"foreignKey:CategoryID"`
}
