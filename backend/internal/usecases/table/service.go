package table

import "github.com/ivan/cafe_reservation/internal/entities"

type service struct {
	tableRepository Repository
}

// CreateTable implements UseCase.
func (self *service) CreateTable(number string, seats int) error {
	return self.tableRepository.CreateTable(number, seats)
}

// DeleteTableById implements UseCase.
func (self *service) DeleteTableById(tableID uint) error {
	return self.tableRepository.DeleteTableById(tableID)
}

// GetTableById implements UseCase.
func (self *service) GetTableById(id string) (*entities.TableEntity, error) {
	return self.tableRepository.GetById(id)
}

// GetTables implements UseCase.
func (self *service) GetAvailableTables() (*[]entities.TableEntityWithAvailable, error) {
	return self.tableRepository.GetAllAvailableTables()
}

// UpdateTableById implements UseCase.
func (self *service) UpdateTableById(tableID uint64, number string, seats int) error {
	return self.tableRepository.UpdateTableById(tableID, number, seats)
}

func NewTableService(r Repository) UseCase {
	return &service{tableRepository: r}
}
