package repository

import (
	"github.com/ivan/cafe_reservation/internal/entities"
	"github.com/ivan/cafe_reservation/internal/usecases/table"
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	Number      string `gorm:"not null"`
	Seats       int    `gorm:"not null;type:integer"`
	IsAvailable bool   `gorm:"-"`
}

func (Table) TableName() string {
	return "tables"
}

type tableRepository struct {
	DB *gorm.DB
}

func (self *tableRepository) UpdateTableById(tableID uint64, number string, seats int) error {
	var table Table

	if err := self.DB.First(&table, tableID).Error; err != nil {
		return err
	}

	table.Number = number
	table.Seats = seats

	if err := self.DB.Save(&table).Error; err != nil {
		return err
	}

	return nil
}

func (self *tableRepository) CreateTable(number string, seats int) error {
	var table Table
	table.Number = number
	table.Seats = seats
	if err := self.DB.Model(&Table{}).Create(&table).Error; err != nil {
		return err
	}
	return nil
}

func (self *tableRepository) DeleteTableById(tableID uint) error {
	if err := self.DB.Model(&Table{}).Where("id = ?", tableID).Delete(&Table{}).Error; err != nil {
		return err
	}
	return nil
}

func (self *tableRepository) GetById(id string) (*entities.TableEntity, error) {
	var table entities.TableEntity
	if err := self.DB.Model(&Table{}).Where("id=?", id).First(&table).Error; err != nil {
		return nil, err
	}
	return &table, nil

}

func (self *tableRepository) GetAllAvailableTables() (*[]entities.TableEntityWithAvailable, error) {
	var tables []entities.TableEntityWithAvailable
	if err := self.DB.Model(&Table{}).Find(&tables).Error; err != nil {
		return nil, err
	}

	for i, table := range tables {
		var count int64
		if err := self.DB.Model(&Reservation{}).
			Where("table_id = ?", table.ID).
			Where("status = ?", "active").
			Where("reserve_time + interval '1 minute' * duration > NOW()").
			Count(&count).Error; err != nil {
			return nil, err
		}

		tables[i].IsAvailable = (count == 0)
	}
	return &tables, nil

}

func NewTableRepository(db *gorm.DB) table.Repository {
	return &tableRepository{DB: db}
}
