package service

import (
	"context"
	"test-golang/data/request"
	"test-golang/data/response"
	"test-golang/helper"
	"test-golang/model"
	"test-golang/repository"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomersRepository
}

func NewCustomerServiceImpl(customer repository.CustomersRepository) CustomerService {
	return &CustomerServiceImpl{CustomerRepository: customer}
}

// Delete implements CustomerService.
func (c *CustomerServiceImpl) Delete(ctx context.Context, id int) {
	customer, err := c.CustomerRepository.ReadById(ctx, id)
	helper.PanicIfError(err)
	c.CustomerRepository.Delete(ctx, customer.Id)
}

// Insert implements CustomerService.
func (c *CustomerServiceImpl) Insert(ctx context.Context, request request.CustomerCreateRequest) {
	customer := model.Customers{
		Id:   request.Id,
		Name: request.Name,
		Age:  request.Age,
	}
	c.CustomerRepository.Insert(ctx, customer)
}

// Read implements CustomerService.
func (c *CustomerServiceImpl) Read(ctx context.Context) ([]response.CustomerResponse, error) {
	customers, err := c.CustomerRepository.Read(ctx)
	helper.PanicIfError(err)

	cusResp := []response.CustomerResponse{}
	for _, v := range customers {
		cusResp = append(cusResp, response.CustomerResponse{
			Id:   v.Id,
			Name: v.Name,
			Age:  v.Age,
		})
	}
	return cusResp, nil
}

// ReadById implements CustomerService.
func (c *CustomerServiceImpl) ReadById(ctx context.Context, id int) (*response.CustomerResponse, error) {
	customer, err := c.CustomerRepository.ReadById(ctx, id)
	if err != nil {
		return nil, err
	}
	cusResp := response.CustomerResponse{
		Id:   customer.Id,
		Name: customer.Name,
		Age:  customer.Age,
	}
	return &cusResp, nil
}

// Update implements CustomerService.
func (c *CustomerServiceImpl) Update(ctx context.Context, request request.CustomerUpdateRequest) error {
	customer, err := c.CustomerRepository.ReadById(ctx, request.Id)
	if err != nil {
		return err
	}
	customer.Age = request.Age
	customer.Name = request.Name
	c.CustomerRepository.Update(ctx, *customer)
	return nil
}
