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
	r.HandleFunc("/product", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/product/{id}", handlers.GetProductById).Methods("GET")
	r.HandleFunc("/order", handlers.CreateOrder).Methods("POST")
	r.HandleFunc("/order/{id}", handlers.GetOrderById).Methods("GET")
	r.HandleFunc("/order-detail", handlers.CreateOderDetail).Methods("POST")
	r.HandleFunc("/order-detail/{id}", handlers.GetOderDetailById).Methods("GET")
	http.ListenAndServe(":8080", r)
}
