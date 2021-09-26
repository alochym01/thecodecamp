package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alochym01/thecodecamp/service"
	"github.com/gorilla/mux"
)

// CustomerHandler ...
type CustomerHandler struct {
	Service service.CustomerService
}

// GetAllCustomers ...
func (c CustomerHandler) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	cs, err := c.Service.GetAllCustomer()

	if err != nil {
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(err.Code)
		// // fmt.Fprintf(w, err.Message)
		// json.NewEncoder(w).Encode(err)
		responseHelper(w, err.Code, cs)
		return
	}

	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(cs)
	responseHelper(w, http.StatusOK, cs)
}

// GetCustomerByID ...
func (c CustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["customer_id"]

	cs, err := c.Service.GetCustomerByID(id)

	if err != nil {
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(err.Code)
		// // fmt.Fprintf(w, err.Message)
		// json.NewEncoder(w).Encode(err)
		responseHelper(w, err.Code, err)
		return
	}

	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(cs)
	responseHelper(w, http.StatusOK, cs)
}

func responseHelper(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
