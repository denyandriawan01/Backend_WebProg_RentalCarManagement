package models

type Rental struct {
	RentalID    int64  `gorm:"primaryKey" json:"rental_id"`
	CarID       int64  `json:"car_id"`
	NIK         string `json:"nik"`
	Name        string `json:"name"`
	UsageRegion string `json:"usage_region"`
	RentalDate  string `json:"rental_date"`
	ReturnDate  string `json:"return_date"`
	TotalPrice  int64  `json:"total_price"`
	IsCompleted bool   `json:"is_completed"`
	Car         Car    `gorm:"references:CarID"`
}
