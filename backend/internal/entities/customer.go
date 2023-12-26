package entities

import "github.com/asaskevich/govalidator"

type CustomerEntity struct {
	ID     uint   `json:"id"`
	UserID uint64 `json:"user_id" valid:"required~User ID is required"`
	Name   string `json:"name" valid:"required~Name is required,alphaNum~Name must contain letters and numbers"`
	Phone  uint64 `json:"phone" valid:"required~Phone number is required"`
	Email  string `json:"email" valid:"required~Email is required,email~Email must be a valid email address"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (c CustomerEntity) Validate() (bool, error) {
	return govalidator.ValidateStruct(c)
}
