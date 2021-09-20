package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alochym01/thecodecamp/service"
)

// CustomerHandler ...
type CustomerHandler struct {
	Service service.CustomerService
}

func (c CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	cs, _ := c.Service.GetAllCustomer()
	json.NewEncoder(w).Encode(cs)
}
