package models

type Category struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex;not null"`

	Offers              []Offer    `gorm:"foreignKey:CategoryId"`
	PreferredCategories []Category `gorm:"many2many:user_categories;"`
}
