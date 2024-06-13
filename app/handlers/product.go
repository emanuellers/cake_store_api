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

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	product := model.Product{}
	body := r.Body
	jsonBody, err := util.JSONDecoderProduct(body, product)

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

	stmt, err := conn.Prepare("INSERT INTO products (name, description, is_available, price, qtd_stored) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(jsonBody.Name, jsonBody.Description, jsonBody.IsAvailable, jsonBody.Price, jsonBody.QtdStored)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Inserido com sucesso!"))

}

func GetProductById(w http.ResponseWriter, r *http.Request) {

	variables := mux.Vars(r)
	id, hasId := variables["id"]

	if !hasId {
		http.Error(w, "Id Producte vazio.", http.StatusBadRequest)
		return
	}

	db := database.DB{}
	conn, err := db.Connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	stmt, err := conn.Prepare("SELECT * FROM products WHERE id = (?)")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	product := model.Product{}
	err = stmt.QueryRow(id).Scan(&product.Name, &product.Description, &product.IsAvailable,
		&product.Price, &product.QtdStored, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	productJson, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(string(productJson)))

}
