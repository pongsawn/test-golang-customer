package router

import (
	"fmt"
	"net/http"
	"test-golang/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(customerController *controller.CustomerController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, `Welcome`)
	})
	router.POST("/customers/create", customerController.Insert)
	router.PUT("/customers/update/:id", customerController.Update)
	router.DELETE("/customers/delete/:id", customerController.Delete)
	router.GET("/customers/read", customerController.Read)
	return router
}
