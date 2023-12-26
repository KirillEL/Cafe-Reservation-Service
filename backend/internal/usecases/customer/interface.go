package customer

import "github.com/ivan/cafe_reservation/internal/entities"

type Reader interface {
	GetAllCustomers() (*[]entities.CustomerEntity, error)
}

type Writer interface {
	DeleteCustomerById(customerID uint) error
	CreateCustomer(userID uint, name string, phone int, email string) error
	UpdateCustomerById(customerID uint, userID uint, name string, phone int, email string) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetAllCustomers() (*[]entities.CustomerEntity, error)
	DeleteCustomerById(customerID uint) error
	CreateCustomer(userID uint, name string, phone int, email string) error
	UpdateCustomerById(customerID uint, userID uint, name string, phone int, email string) error
}
