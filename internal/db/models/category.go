package models

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex;not null"`

	Offers []Offer `gorm:"foreignKey:CategoryID"`
}
