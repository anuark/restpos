package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Order .
type Order struct {
	Model
	CustomerID uint `json:"customer_id"`
}

// OrderIndex .
func OrderIndex(w http.ResponseWriter, r *http.Request) {
	var orders []Order
	Db.Find(&orders)

	ordersJSON, err := json.Marshal(orders)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(ordersJSON))
}

// OrderCreate .
func OrderCreate(w http.ResponseWriter, r *http.Request) {
	var order Order
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&order)

	Db.Create(&order)

	orderJSON, err := json.Marshal(order)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(orderJSON))
}

// OrderUpdate .
func OrderUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var order Order
	Db.First(&order, id)
	if order.ID == 0 {
		http.Error(w, "order record not found.", http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&order)

	Db.Save(&order)

	orderJSON, err := json.Marshal(order)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(orderJSON))
}

// OrderView .
func OrderView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var order Order
	Db.First(&order, id)
	if order.ID == 0 {
		http.Error(w, "order record not found.", http.StatusNotFound)
		return
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(orderJSON))
}

// OrderDelete .
func OrderDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var order Order
	Db.First(&order, id)
	if order.ID == 0 {
		http.Error(w, "order record not found.", http.StatusNotFound)
		return
	}

	Db.Delete(&order)

	res, _ := json.Marshal(struct {
		Data string `json:"data"`
	}{"success"})

	fmt.Fprint(w, res)
}
