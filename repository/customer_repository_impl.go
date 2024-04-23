package repository

import (
	"context"
	"database/sql"
	"errors"
	"test-golang/helper"
	"test-golang/model"
)

type CustomerRepositoryImpl struct {
	Db *sql.DB
}

func NewCustomerRepository(Db *sql.DB) CustomersRepository {
	return &CustomerRepositoryImpl{Db: Db}
}

// Delete implements CustomersRepository.
func (c *CustomerRepositoryImpl) Delete(ctx context.Context, id int) {
	tx, err := c.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := `DELETE FROM customers WHERE id = $1`
	_, err = tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}

// Insert implements CustomersRepository.
func (c *CustomerRepositoryImpl) Insert(ctx context.Context, customer model.Customers) {
	tx, err := c.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := `INSERT INTO customers (id,name,age) VALUES ($1,$2,$3)`
	_, err = tx.ExecContext(ctx, SQL, customer.Id, customer.Name, customer.Age)
	helper.PanicIfError(err)
}

// Read implements CustomersRepository.
func (c *CustomerRepositoryImpl) ReadById(ctx context.Context, id int) (*model.Customers, error) {
	tx, err := c.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := `SELECT * FROM customers WHERE id = $1`
	row, err := tx.Query(SQL, id)
	helper.PanicIfError(err)
	defer row.Close()

	customer := model.Customers{}
	if row.Next() {
		err := row.Scan(&customer.Id, &customer.Name, &customer.Age)
		helper.PanicIfError(err)
		return &customer, nil
	} else {
		return nil, errors.New(`customer id not found`)
	}
}

// Update implements CustomersRepository.
func (c *CustomerRepositoryImpl) Update(ctx context.Context, customer model.Customers) {
	tx, err := c.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := `UPDATE customers set name = $1, age = $2 where id = $3`
	_, err = tx.ExecContext(ctx, SQL, customer.Name, customer.Age, customer.Id)
	helper.PanicIfError(err)
}

// Read implements CustomersRepository.
func (c *CustomerRepositoryImpl) Read(ctx context.Context) ([]model.Customers, error) {
	tx, err := c.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := `SELECT * FROM customers`
	row, err := tx.Query(SQL)
	helper.PanicIfError(err)
	customers := []model.Customers{}
	for row.Next() {
		customer := model.Customers{}
		err := row.Scan(&customer.Id, &customer.Name, &customer.Age)
		helper.PanicIfError(err)
		customers = append(customers, customer)
	}
	return customers, nil
}
