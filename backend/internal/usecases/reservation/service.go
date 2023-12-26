package reservation

import (
	"time"

	"github.com/ivan/cafe_reservation/internal/entities"
)

type service struct {
	reservationRepository Repository
}

// UpdateReservationById implements UseCase.
func (self *service) UpdateReservationById(ReservationID uint, CustomerID uint, TableID uint, ReserveTime time.Time, Duration int, Status string) error {
	return self.reservationRepository.UpdateReservationById(ReservationID, CustomerID, TableID, ReserveTime, Duration, Status)
}

// CreateReservationByAdmin implements UseCase.
func (self *service) CreateReservationByAdmin(customerID uint, TableID uint, reserve_time string, duration int, status string) error {
	return self.reservationRepository.CreateReservationByAdmin(customerID, TableID, reserve_time, duration, status)
}

// GetAllReservations implements UseCase.
func (self *service) GetAllReservations() (*[]entities.ReservationEntity, error) {
	return self.reservationRepository.GetAllReservations()
}

// GetReservationsByUserId implements UseCase.
func (self *service) GetReservationsByUserId(userID uint) (*[]entities.ReservationsDetails, error) {
	return self.reservationRepository.GetReservationsByUserId(userID)
}

// CreateReservation implements UseCase.
func (self *service) CreateReservation(userID uint, name string, phone string, email string, tableNumber string, reserveTime string, duration string) error {
	return self.reservationRepository.Create(userID, name, phone, email, tableNumber, reserveTime, duration)
}

// DeleteReservationById implements UseCase.
func (self *service) DeleteReservationById(id int) (bool, error) {
	return self.reservationRepository.Delete(id)
}

func NewReservationService(r Repository) UseCase {
	return &service{reservationRepository: r}
}
