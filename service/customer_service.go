package service

import (
	"context"
	"test-golang/data/request"
	"test-golang/data/response"
)

type CustomerService interface {
	Insert(ctx context.Context, request request.CustomerCreateRequest)
	Update(ctx context.Context, request request.CustomerUpdateRequest) error
	Delete(ctx context.Context, id int)
	ReadById(ctx context.Context, id int) (*response.CustomerResponse, error)
	Read(ctx context.Context) ([]response.CustomerResponse, error)
}
