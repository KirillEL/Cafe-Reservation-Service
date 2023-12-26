package reservation

import (
	"time"

	"github.com/ivan/cafe_reservation/internal/entities"
)

type Reader interface {
	GetReservationsByUserId(userID uint) (*[]entities.ReservationsDetails, error)
	GetAllReservations() (*[]entities.ReservationEntity, error)
}

type Writer interface {
	Create(userID uint, name string, phone string, email string, tableNumber string, reserveTime string, duration string) error
	Delete(id int) (bool, error)
	CreateReservationByAdmin(customerID uint, TableID uint, reserve_time string, duration int, status string) error
	UpdateReservationById(ReservationID uint, CustomerID uint, TableID uint, ReserveTime time.Time, Duration int, Status string) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateReservation(userID uint, name string, phone string, email string, tableNumber string, reserveTime string, duration string) error
	CreateReservationByAdmin(customerID uint, TableID uint, reserve_time string, duration int, status string) error
	DeleteReservationById(id int) (bool, error)
	UpdateReservationById(ReservationID uint, CustomerID uint, TableID uint, ReserveTime time.Time, Duration int, Status string) error
	GetReservationsByUserId(userID uint) (*[]entities.ReservationsDetails, error)
	GetAllReservations() (*[]entities.ReservationEntity, error)
}
