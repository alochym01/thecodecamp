package app

// import (
// 	"log"
// 	"net/http"

// 	"github.com/alochym01/thecodecamp/domain"
// 	"github.com/alochym01/thecodecamp/service"
// 	"github.com/gorilla/mux"
// )

// // Start ...
// func Start() {
// 	router := mux.NewRouter()

// 	cHandler := CustomerHandler{service: service.NewCustomerService(domain.NewUserRepositoryStubs())}
// 	router.HandleFunc("/customer", cHandler.getAllCustomers).Methods(http.MethodGet)

// 	log.Fatal(http.ListenAndServe(":8000", router))
// }
