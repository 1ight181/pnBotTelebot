package enums

type OfferStatus string

// Реализация через обычные строки, а не через iota ввиду работы с gorm
// который не умеет сериализовать int-ы iota в строки
const (
	StatusActive   OfferStatus = "active"
	StatusInactive OfferStatus = "inactive"
	StatusArchived OfferStatus = "archived"
)
