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

type ClientHandler struct{}

func CreateClient(w http.ResponseWriter, r *http.Request) {

	client := model.Client{}
	body := r.Body
	jsonBody, err := util.JSONDecoderClient(body, client)

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

	}
	defer conn.Close()

	stmt, err := conn.Prepare("INSERT INTO clients (firstname, lastname, email) VALUES (?, ?, ?)")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer stmt.Close()

	_, err = stmt.Exec(jsonBody.FirstName, jsonBody.LastName, jsonBody.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(200)
	w.Write([]byte("Inserido com sucesso!"))

}

func GetClientById(w http.ResponseWriter, r *http.Request) {

	variables := mux.Vars(r)
	id, hasId := variables["id"]

	if !hasId {
		http.Error(w, "Id cliente vazio.", http.StatusBadRequest)
		return
	}

	db := database.DB{}
	conn, err := db.Connect()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	defer conn.Close()

	stmt, err := conn.Prepare("SELECT * FROM clients WHERE id = (?)")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer stmt.Close()

	client := model.Client{}
	err = stmt.QueryRow(id).Scan(&client.FirstName, &client.LastName, &client.Email, &client.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	clientJson, err := json.Marshal(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(200)
	w.Write([]byte(string(clientJson)))

}
