package entities

import "github.com/asaskevich/govalidator"

type TableEntity struct {
	Number string `json:"number" valid:"required~Table number is required,alphanum~Table number must be alphanumeric"`
	Seats  int    `json:"seats" valid:"required~Number of seats is required,range(1|10)~Seats must be between 1 and 10"`
}

type TableEntityWithAvailable struct {
	ID          uint64 `json:"id"`
	Number      string `json:"number"`
	Seats       uint   `json:"seats"`
	IsAvailable bool   `json:"is_available" gorm:"-"`
}

// add govalidator
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Validate performs the validation on TableEntity fields based on the tags.
func (table TableEntity) Validate() (bool, error) {
	return govalidator.ValidateStruct(table)
}
