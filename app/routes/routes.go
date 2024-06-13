package route

import (
	"net/http"

	"github.com/emanuellers/cake_store_api/app/handlers"
	"github.com/gorilla/mux"
)

func Routes() {
	r := mux.NewRouter()
	r.HandleFunc("/client/{id}", handlers.GetClientById).Methods("GET")
	r.HandleFunc("/client", handlers.CreateClient).Methods("POST")
	http.ListenAndServe(":8080", r)
}
