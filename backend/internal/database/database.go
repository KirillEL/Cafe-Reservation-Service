package database

import (
	"log"

	"github.com/ivan/cafe_reservation/internal/repository"
	"gorm.io/gorm"
)

func ConnectToDB(dialector gorm.Dialector) (*gorm.DB, error) {
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := gormDB.Exec("Create TYPE role_type as ENUM('admin', 'user');").Error; err != nil {
		log.Println(err)
	}

	if err := gormDB.Exec("CREATE TYPE reservation_type AS ENUM ('active', 'decline', 'ended');").Error; err != nil {
		log.Println(err)
	}

	gormDB.AutoMigrate(&repository.User{}, &repository.Customer{}, &repository.Reservation{}, &repository.Table{})

	if err := gormDB.Exec("INSERT INTO users(login, email, password, role) values('admin','admin@admin.com', 'admin1', 'admin');").Error; err != nil {
		log.Println(err)
	}

	return gormDB, nil
}
