package app

import (
	"encoding/json"
	"net/http"

	"github.com/alochym01/thecodecamp/service"
)

// CustomerHandler ...
type CustomerHandler struct {
	service service.CustomerService
}

func (c CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	cs, _ := c.service.GetAllCustomer()
	json.NewEncoder(w).Encode(cs)
}
