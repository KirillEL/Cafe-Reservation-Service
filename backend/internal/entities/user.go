package entities

import (
	"github.com/asaskevich/govalidator"
	"github.com/ivan/cafe_reservation/internal/types"
)

type UserEntity struct {
	ID    uint       `json:"id" valid:"required~id is required"`
	Login string     `json:"login" valid:"required~Login is required"`
	Email string     `json:"email" valid:"required~Email is required"`
	Role  types.Role `json:"role" valid:"required~Role is required"`
}

type AllUsersEntity struct {
	ID       uint   `json:"id" valid:"required~id is required"`
	Login    string `json:"login" valid:"required~Login is required"`
	Password string `json:"password" valid:"required~Password is required"`
	Email    string `json:"email" valid:"required~Email is required"`
	Role     string `json:"role" valid:"required~Role is required"`
}

type UserCreateEntity struct {
	Login    string     `json:"login"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Role     types.Role `json:"role"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (user UserEntity) Validate() (bool, error) {
	return govalidator.ValidateStruct(user)
}
