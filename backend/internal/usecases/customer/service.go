package customer

import "github.com/ivan/cafe_reservation/internal/entities"

type service struct {
	customerRepository Repository
}

func (self *service) UpdateCustomerById(customerID uint, userID uint, name string, phone int, email string) error {
	return self.customerRepository.UpdateCustomerById(customerID, userID, name, phone, email)
}

func (self *service) CreateCustomer(userID uint, name string, phone int, email string) error {
	return self.customerRepository.CreateCustomer(userID, name, phone, email)
}

func (self *service) DeleteCustomerById(customerID uint) error {
	return self.customerRepository.DeleteCustomerById(customerID)
}

func (self *service) GetAllCustomers() (*[]entities.CustomerEntity, error) {
	return self.customerRepository.GetAllCustomers()
}

func NewCustomerService(r Repository) UseCase {
	return &service{customerRepository: r}
}
