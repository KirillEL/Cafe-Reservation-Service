package table

import "github.com/ivan/cafe_reservation/internal/entities"

type Reader interface {
	GetAllAvailableTables() (*[]entities.TableEntityWithAvailable, error)
	GetById(id string) (*entities.TableEntity, error)
}

type Writer interface {
	DeleteTableById(tableID uint) error
	CreateTable(number string, seats int) error
	UpdateTableById(tableID uint64, number string, seats int) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateTable(number string, seats int) error
	GetTableById(id string) (*entities.TableEntity, error)
	GetAvailableTables() (*[]entities.TableEntityWithAvailable, error)
	UpdateTableById(tableID uint64, number string, seats int) error
	DeleteTableById(tableID uint) error
}
