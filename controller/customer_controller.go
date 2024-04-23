package controller

import (
	"net/http"
	"strconv"
	"test-golang/data/request"
	"test-golang/data/response"
	"test-golang/helper"
	"test-golang/service"

	"github.com/julienschmidt/httprouter"
)

type CustomerController struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerServic service.CustomerService) *CustomerController {
	return &CustomerController{CustomerService: customerServic}
}

func (controller *CustomerController) Insert(writer http.ResponseWriter, requests *http.Request, param httprouter.Params) {
	customerCreateRequest := request.CustomerCreateRequest{}
	helper.ReadRequestBody(requests, &customerCreateRequest)

	controller.CustomerService.Insert(requests.Context(), customerCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CustomerController) Update(writer http.ResponseWriter, requests *http.Request, param httprouter.Params) {
	customerUpdateRequest := request.CustomerUpdateRequest{}
	helper.ReadRequestBody(requests, &customerUpdateRequest)

	cusId := param.ByName(`id`)
	id, err := strconv.Atoi(cusId)
	helper.PanicIfError(err)
	customerUpdateRequest.Id = id
	//เช็คก่อนว่ามี id ในระบบไหม
	_, err = controller.CustomerService.ReadById(requests.Context(), id)
	if err != nil {
		webResponse := response.WebResponse{
			Code:   404,
			Status: "customer id not found",
			Data:   nil,
		}
		helper.WriteResponseBody(writer, webResponse)
	} else {
		controller.CustomerService.Update(requests.Context(), customerUpdateRequest)
		webResponse := response.WebResponse{
			Code:   200,
			Status: "Ok",
			Data:   nil,
		}
		helper.WriteResponseBody(writer, webResponse)
	}
}

func (controller *CustomerController) Delete(writer http.ResponseWriter, requests *http.Request, param httprouter.Params) {
	cusId := param.ByName(`id`)
	id, err := strconv.Atoi(cusId)
	helper.PanicIfError(err)

	controller.CustomerService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CustomerController) ReadById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	cusId := param.ByName(`id`)
	id, err := strconv.Atoi(cusId)
	helper.PanicIfError(err)

	result, err := controller.CustomerService.ReadById(request.Context(), id)
	helper.PanicIfError(err)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CustomerController) Read(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	result, err := controller.CustomerService.Read(request.Context())
	helper.PanicIfError(err)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}
