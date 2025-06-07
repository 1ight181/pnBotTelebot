package models

type Category struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex;not null"`

	Offers []Offer `gorm:"foreignKey:CategoryId"`
}
