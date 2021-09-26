package router

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/alochym01/thecodecamp/domain"
	"github.com/alochym01/thecodecamp/handlers"
	"github.com/alochym01/thecodecamp/service"
	"github.com/gorilla/mux"
)

// Start ...
func Start(db *sql.DB) {
	router := mux.NewRouter()

	// Memory storage
	// cHandler := handlers.CustomerHandler{
	// 	Service: service.NewCustomerService(domain.NewUserRepositoryStubs()),
	// }

	// DB storage
	cHandler := handlers.CustomerHandler{
		Service: service.NewCustomerService(domain.NewUserRepositoryDB(db)),
	}
	router.HandleFunc("/customer", cHandler.GetAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", cHandler.GetCustomerByID).Methods(http.MethodGet)

	fmt.Println("App start 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
