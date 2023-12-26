package repository

import (
	"errors"

	"github.com/ivan/cafe_reservation/internal/entities"
	"github.com/ivan/cafe_reservation/internal/usecases/customer"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	UserID uint64 `gorm:"type:bigint;not null"`
	Name   string `gorm:"not null"`
	Phone  uint64 `gorm:"type:bigint;not null"`
	Email  string `gorm:"type:text"`

	User User `gorm:"foreignKey:UserID"`
}

func (Customer) TableName() string {
	return "customers"
}

type customerRepository struct {
	DB *gorm.DB
}

// UpdateCustomerById implements customer.Repository.
func (self *customerRepository) UpdateCustomerById(customerID uint, userID uint, name string, phone int, email string) error {
	var customer entities.CustomerEntity
	result := self.DB.Model(&Customer{}).First(&customer, customerID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return result.Error
	}
	customer.UserID = uint64(userID)
	customer.Name = name
	customer.Phone = uint64(phone)
	customer.Email = email

	if err := self.DB.Table("customers").Where("id = ?", customerID).Save(&customer).Error; err != nil {
		return err
	}
	return nil
}

// CreateCustomer implements customer.Repository.
func (self *customerRepository) CreateCustomer(userID uint, name string, phone int, email string) error {
	var customer Customer
	customer.UserID = uint64(userID)
	customer.Name = name
	customer.Phone = uint64(phone)
	customer.Email = email

	if err := self.DB.Model(&Customer{}).Create(&customer).Error; err != nil {
		return err
	}
	return nil
}

// DeleteCustomerById implements customer.Repository.
func (self *customerRepository) DeleteCustomerById(customerID uint) error {
	if err := self.DB.Model(&Customer{}).Where("id = ?", customerID).Delete(&Customer{}); err != nil {
		return err.Error
	}
	return nil
}

// GetAllCustomers implements customer.Repository.
func (self *customerRepository) GetAllCustomers() (*[]entities.CustomerEntity, error) {
	var customers []entities.CustomerEntity
	if err := self.DB.Model(&Customer{}).Find(&customers).Error; err != nil {
		return nil, err
	}
	return &customers, nil
}

func NewCustomerRepository(db *gorm.DB) customer.Repository {
	return &customerRepository{DB: db}
}
