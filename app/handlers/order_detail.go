package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emanuellers/cake_store_api/app/util"
	"github.com/emanuellers/cake_store_api/database"
	"github.com/emanuellers/cake_store_api/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func CreateOderDetail(w http.ResponseWriter, r *http.Request) {

	orderDetail := model.OderDetail{}
	body := r.Body
	jsonBody, err := util.JSONDecoderOderDetail(body, orderDetail)

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

	stmt, err := conn.Prepare("INSERT INTO order_details (order_id, product_id, product_name, product_price) VALUES (?, ?, ?, ?)")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(jsonBody.OrderId, jsonBody.ProductId, jsonBody.ProductName, jsonBody.ProductPrice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Inserido com sucesso!"))

}

func GetOderDetailById(w http.ResponseWriter, r *http.Request) {

	variables := mux.Vars(r)
	id, hasId := variables["id"]

	if !hasId {
		http.Error(w, "Id OderDetaile vazio.", http.StatusBadRequest)
		return
	}

	db := database.DB{}
	conn, err := db.Connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	stmt, err := conn.Prepare("SELECT * FROM orderDetails WHERE id = (?)")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	orderDetail := model.OderDetail{}
	err = stmt.QueryRow(id).Scan(&orderDetail.OrderId, &orderDetail.ProductId, &orderDetail.ProductName, &orderDetail.ProductPrice, &orderDetail.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orderDetailJson, err := json.Marshal(orderDetail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(string(orderDetailJson)))

}
