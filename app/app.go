package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start func
func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/greet", greeting).Methods(http.MethodGet)
	router.HandleFunc("/customer", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post method")
}
