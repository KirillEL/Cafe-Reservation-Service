package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type ReservationStatus string

const (
	ActiveStatus   ReservationStatus = "active"
	DeclinedStatus ReservationStatus = "decline"
	EndedStatus    ReservationStatus = "ended"
)

type ReservationEntity struct {
	ID          uint64            `json:"id"`
	CustomerID  uint64            `json:"customer_id"`
	TableID     uint64            `json:"table_id"`
	ReserveTime time.Time         `json:"reserve_time"`
	Duration    int               `json:"duration"`
	Status      ReservationStatus `json:"status"`
}

type ReservationsDetails struct {
	ID          uint64            `json:"id"`
	ReserveTime time.Time         `json:"reserve_time"`
	Duration    int               `json:"duration"`
	Status      ReservationStatus `json:"status"`
	Name        string            `json:"name"`
	Phone       uint64            `json:"phone"`
	Email       string            `json:"email"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (r ReservationEntity) Validate() (bool, error) {
	return govalidator.ValidateStruct(r)
}
