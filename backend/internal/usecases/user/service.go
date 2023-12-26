package user

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ivan/cafe_reservation/internal/entities"
	"github.com/ivan/cafe_reservation/internal/types"
)

type service struct {
	userRepository Repository
}

// CreateUser implements UseCase.
func (self *service) CreateUser(login string, email string, password string, role types.Role) error {
	return self.userRepository.CreateUser(login, email, password, role)
}

// DeleteUserById implements UseCase.
func (self *service) DeleteUserById(userID uint) error {
	return self.userRepository.DeleteUserById(userID)
}

// GetAllUsers implements UseCase.
func (self *service) GetAllUsers() (*[]entities.AllUsersEntity, error) {
	return self.userRepository.GetAllUsers()
}

// UpdateUserById implements UseCase.
func (self *service) UpdateUserById(userID uint, login string, email string, password string, role types.Role) error {
	return self.userRepository.UpdateUserById(userID, login, email, password, role)
}

// Login implements UseCase.
func (s *service) Login(email string, password string) (*entities.UserEntity, error) {
	return s.userRepository.LoginUser(email, password)
}

// Register implements UseCase.
func (s *service) Register(login string, email string, password string) (*entities.UserEntity, error) {
	user, err := s.userRepository.RegisterUser(login, email, password)
	log.Println(user)
	log.Println(err)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Verify implements UseCase.
func (s *service) Verify(authToken string) (*jwt.Token, error) {
	panic("unimplemented")
}

func NewUserService(r Repository) UseCase {
	return &service{userRepository: r}
}
