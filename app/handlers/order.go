package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/emanuellers/cake_store_api/app/util"
	"github.com/emanuellers/cake_store_api/database"
	"github.com/emanuellers/cake_store_api/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {

	order := model.Order{}
	body := r.Body
	jsonBody, err := util.JSONDecoderOrder(body, order)

	// out, _ := json.Marshal(jsonBody)
	// http.Error(w, string(out), http.StatusInternalServerError)
	// return

	// http.Error(w, a, http.StatusInternalServerError)
	// return

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db := database.DB{}
	conn, err := db.Connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	defer conn.Close()

	stmt, err := conn.Prepare("INSERT INTO orders (client_id, ordered_at, total_price, delivery_date, description) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	// deliveryAt, err := time.Parse(time.RFC3339, jsonBody.DeliveryDate.String())

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	_, err = stmt.Exec(jsonBody.ClientId, time.Now(), jsonBody.TotalPrice, jsonBody.DeliveryDate, jsonBody.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Inserido com sucesso!"))

}

func GetOrderById(w http.ResponseWriter, r *http.Request) {

	variables := mux.Vars(r)
	id, hasId := variables["id"]

	if !hasId {
		http.Error(w, "Id Ordere vazio.", http.StatusBadRequest)
		return
	}

	db := database.DB{}
	conn, err := db.Connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	stmt, err := conn.Prepare("SELECT * FROM orders WHERE id = (?)")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	order := model.Order{}
	err = stmt.QueryRow(id).Scan(&order.ClientId, &order.OrderedAt, &order.TotalPrice, &order.DeliveryDate, &order.Description, &order.UpdatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orderJson, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(string(orderJson)))

}
