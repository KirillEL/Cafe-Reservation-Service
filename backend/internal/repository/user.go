package repository

import (
	"errors"

	"github.com/ivan/cafe_reservation/internal/entities"
	"github.com/ivan/cafe_reservation/internal/types"
	"github.com/ivan/cafe_reservation/internal/usecases/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string     `gorm:"type:varchar(50);not null;unique"`
	Email    string     `gorm:"type:text;not null`
	Password string     `gorm:"type:varchar(100);not null`
	Role     types.Role `gorm:"type:role_type;not null;default:'user'"`
}

func (User) TableName() string {
	return "users"
}

type userRepository struct {
	DB *gorm.DB
}

func (self *userRepository) CreateUser(login string, email string, password string, role types.Role) error {
	userEntity := User{
		Login:    login,
		Email:    email,
		Password: password,
		Role:     role,
	}
	result := self.DB.Model(&User{}).Create(&userEntity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (self *userRepository) DeleteUserById(userID uint) error {
	result := self.DB.Model(&User{}).Where("id = ?", userID).Delete(&User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (self *userRepository) GetAllUsers() (*[]entities.AllUsersEntity, error) {
	var users []entities.AllUsersEntity
	if err := self.DB.Model(&User{}).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (self *userRepository) UpdateUserById(userID uint, login string, email string, password string, role types.Role) error {

	user := entities.UserCreateEntity{}
	result := self.DB.Model(&User{}).First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return result.Error
	}

	user.Login = login
	user.Email = email
	user.Password = password
	user.Role = role

	result = self.DB.Table("users").Where("id = ?", userID).Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (self *userRepository) LoginUser(email string, password string) (*entities.UserEntity, error) {
	var userDto User
	result := self.DB.Where("email = ? AND password = ?", email, password).First(&userDto)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("Incorrect email or password")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userEntity := &entities.UserEntity{
		ID:    userDto.ID,
		Login: userDto.Login,
		Email: userDto.Email,
		Role:  userDto.Role,
	}
	return userEntity, nil
}

// RegisterUser implements user.Repository.
func (self *userRepository) RegisterUser(login string, email string, password string) (*entities.UserEntity, error) {
	userDto := &User{
		Login:    login,
		Email:    email,
		Password: password,
		Role:     "user",
	}
	if err := self.DB.Model(&User{}).Create(&userDto).Error; err != nil {
		return nil, err
	}
	var entity entities.UserEntity
	entity.Login = userDto.Login
	entity.Email = userDto.Email
	entity.Role = "user"
	return &entity, nil

}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{DB: db}
}
