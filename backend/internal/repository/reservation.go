package repository

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ivan/cafe_reservation/internal/entities"
	"github.com/ivan/cafe_reservation/internal/usecases/reservation"
	"gorm.io/gorm"
)

type ReserveStatus string

const (
	ActiveStatus   ReserveStatus = "active"
	DeclinedStatus ReserveStatus = "decline"
	EndedStatus    ReserveStatus = "ended"
)

type Reservation struct {
	gorm.Model
	CustomerID  uint64        `gorm:"type:bigint;not null"`
	TableID     uint64        `gorm:"type:bigint;not null"`
	ReserveTime time.Time     `gorm:"not null"`
	Duration    int           `gorm:"not null"`
	Status      ReserveStatus `gorm:"type:reservation_type;not null"`

	Customer Customer `gorm:"foreignKey:CustomerID"`
	Table    Table    `gorm:"foreignKey:TableID"`
}

func (Reservation) TableName() string {
	return "reservations"
}

type reservationRepository struct {
	DB *gorm.DB
}

// UpdateReservationById implements reservation.Repository.
func (self *reservationRepository) UpdateReservationById(ReservationID uint, CustomerID uint, TableID uint, ReserveTime time.Time, Duration int, Status string) error {
	var reservation entities.ReservationEntity
	if err := self.DB.Model(&Reservation{}).First(&reservation, ReservationID).Error; err != nil {
		return err
	}
	reservation.CustomerID = uint64(CustomerID)
	reservation.TableID = uint64(TableID)
	reservation.ReserveTime = ReserveTime
	reservation.Duration = Duration
	reservation.Status = entities.ReservationStatus(Status)

	if err := self.DB.Table("reservations").Where("id = ?", ReservationID).Save(&reservation).Error; err != nil {
		return err
	}
	return nil
}

// CreateReservationByAdmin implements reservation.Repository.
func (self *reservationRepository) CreateReservationByAdmin(customerID uint, TableID uint, reserve_time string, duration int, status string) error {
	var reserve Reservation

	parsedTime, err := time.Parse("2006-01-02T15:04", reserve_time)
	if err != nil {
		return err
	}

	reserve.CustomerID = uint64(customerID)
	reserve.TableID = uint64(TableID)
	reserve.ReserveTime = parsedTime
	reserve.Duration = duration
	reserve.Status = ReserveStatus(status)

	if err := self.DB.Model(&Reservation{}).Create(&reserve).Error; err != nil {
		return err
	}
	return nil

}

// GetAllReservations implements reservation.Repository.
func (self *reservationRepository) GetAllReservations() (*[]entities.ReservationEntity, error) {
	var r []entities.ReservationEntity
	result := self.DB.Model(&Reservation{}).Find(&r)
	if result.Error != nil {
		return nil, result.Error
	}
	return &r, nil
}

// GetReservationsByUserId implements reservation.Repository.
func (self *reservationRepository) GetReservationsByUserId(userID uint) (*[]entities.ReservationsDetails, error) {

	var reservations []entities.ReservationsDetails

	err := self.DB.Table("reservations").
		Select("reservations.id, reservations.reserve_time, reservations.duration, reservations.status, customers.name, customers.phone, customers.email").
		Joins("JOIN customers ON customers.id = reservations.customer_id").
		Joins("JOIN users ON users.id = customers.user_id").
		Joins("JOIN tables ON tables.id = reservations.table_id").
		Where("users.id = ?", userID).
		Where("reservations.deleted_at is NULL").
		Where("reservations.status = ?", "active").
		Scan(&reservations).Error

	if err != nil {
		return nil, err
	}
	return &reservations, nil
}

// Delete implements reservation.Repository.
func (self *reservationRepository) Delete(id int) (bool, error) {
	var reservation Reservation

	if err := self.DB.Where("id = ?", id).Delete(&reservation).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (self *reservationRepository) Create(userID uint, name string, phone string, email string, tableNumber string, reserveTime string, duration string) error {
	tx := self.DB.Begin()

	reserveTimeTime, err := time.Parse(time.RFC3339, reserveTime)
	if err != nil {
		tx.Rollback()
		return err
	}

	phoneUint, err := strconv.ParseUint(phone, 10, 64)
	if err != nil {
		tx.Rollback()
		return err
	}

	parts := strings.Split(duration, ":")
	if len(parts) != 2 {
		tx.Rollback()
		return fmt.Errorf("invalid duration format")
	}
	hours, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		tx.Rollback()
		return err
	}
	minutes, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		tx.Rollback()
		return err
	}
	durationInt := int(hours*60 + minutes)

	var table struct {
		ID uint64
	}

	if err := tx.Model(&Table{}).Where("number = ?", tableNumber).First(&table).Error; err != nil {
		tx.Rollback()
		return err
	}

	customer := Customer{
		UserID: uint64(userID),
		Name:   name,
		Phone:  phoneUint,
		Email:  email,
	}
	if err := tx.Create(&customer).Error; err != nil {
		tx.Rollback()
		return err
	}
	reservation := Reservation{
		CustomerID:  uint64(customer.ID),
		TableID:     table.ID,
		ReserveTime: reserveTimeTime,
		Duration:    durationInt,
		Status:      ActiveStatus,
	}

	if err := tx.Create(&reservation).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func NewReservationRepository(db *gorm.DB) reservation.Repository {
	return &reservationRepository{DB: db}
}
