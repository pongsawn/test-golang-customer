package main

import (
	"fmt"
	"net/http"
	"test-golang/config"
	"test-golang/controller"
	"test-golang/helper"
	"test-golang/repository"
	"test-golang/router"
	"test-golang/service"
	"test-golang/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println(`START SERVER`)

	//สร้าง database customers
	db := config.DatabaseConnection()
	statement, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS customers (id INTEGER PRIMARY KEY , name TEXT, age INTEGER)`)
	statement.Exec()

	//สร้าง repository CRUD
	customerRepository := repository.NewCustomerRepository(db)

	//สร้าง service
	customerService := service.NewCustomerServiceImpl(customerRepository)

	//สร้าง controller สำหรับควบคุมการทำงานของ CRUD
	customerController := controller.NewCustomerController(customerService)

	//router ของ CRUD
	routes := router.NewRouter(customerController)
	str := sqlx.Map{
		`name`: `potae`,
	}
	fmt.Printf(`%v`, str.String(`name`))
	//run server
	server := http.Server{Addr: "localhost:8888", Handler: routes}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
