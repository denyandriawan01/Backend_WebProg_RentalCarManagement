package models

type MaintenanceHistory struct {
	MaintenanceID   int64  `gorm:"primaryKey" json:"maintenance_id"`
	CarID           int64  `json:"car_id"`
	MaintenanceDate string `json:"maintenance_date"`
	LastOdometer    int64  `json:"last_odometer"`
	Type            string `json:"type"`
	Description     string `json:"description"`
	Expense         int64  `json:"expense"`
	Car             Car    `gorm:"references:CarID"`
}
