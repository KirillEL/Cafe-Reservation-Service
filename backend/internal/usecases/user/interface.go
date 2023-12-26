package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ivan/cafe_reservation/internal/entities"
	"github.com/ivan/cafe_reservation/internal/types"
)

type Reader interface {
	GetAllUsers() (*[]entities.AllUsersEntity, error)
}

type Writer interface {
	LoginUser(email string, password string) (*entities.UserEntity, error)
	RegisterUser(login string, email string, password string) (*entities.UserEntity, error)
	CreateUser(login string, email string, password string, role types.Role) error
	UpdateUserById(userID uint, login string, email string, password string, role types.Role) error
	DeleteUserById(userID uint) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Login(email string, password string) (*entities.UserEntity, error)
	Register(login string, email string, password string) (*entities.UserEntity, error)
	Verify(authToken string) (*jwt.Token, error)
	CreateUser(login string, email string, password string, role types.Role) error
	UpdateUserById(userID uint, login string, email string, password string, role types.Role) error
	DeleteUserById(userID uint) error
	GetAllUsers() (*[]entities.AllUsersEntity, error)
}
