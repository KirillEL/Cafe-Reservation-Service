package database

import (
	"github.com/ivan/cafe_reservation/internal/repository"
	"gorm.io/gorm"
)

func ConnectToDB(dialector gorm.Dialector) (*gorm.DB, error) {
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	gormDB.AutoMigrate(&repository.User{}, &repository.Customer{}, &repository.Reservation{}, &repository.Table{})

	return gormDB, nil
}
