package repository

import (
	"context"
	"test-golang/model"
)

type CustomersRepository interface {
	Insert(ctx context.Context, customer model.Customers)
	Update(ctx context.Context, customer model.Customers)
	Delete(ctx context.Context, id int)
	ReadById(ctx context.Context, id int) (*model.Customers, error)
	Read(ctx context.Context) ([]model.Customers, error)
}
