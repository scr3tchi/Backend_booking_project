package db

import (
	"github.com/scr3tchi/tiket-booking-project/models"
	"gorm.io/gorm"
)

func DBMigration(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{})
}
