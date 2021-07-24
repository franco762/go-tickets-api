package database

import (
	"franco762/go-tickets-api/internal/ticket"

	"github.com/jinzhu/gorm"
)

// MigrateDB - migra la base de datos y crea la tabla tickets
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&ticket.Ticket{}); result.Error != nil {
		return result.Error
	}
	return nil
}
