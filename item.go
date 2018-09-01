package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Item .
type Item struct {
	gorm.Model
	Name, ImagePath string
	Enabled         bool
	CategoryID      uint
}

// ItemIndex .
func ItemIndex(w http.ResponseWriter, r *http.Request) {
	var items []Item
	Db.Find(&items)

	itemsJSON, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(itemsJSON))
}

// ItemCreate .
func ItemCreate(w http.ResponseWriter, r *http.Request) {
	var item Item
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&item)

	Db.Create(&item)

	itemJSON, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(itemJSON))
}

// ItemUpdate .
func ItemUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var item Item
	Db.First(&item, id)
	if item.ID == 0 {
		http.Error(w, "item record not found.", http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&item)

	Db.Save(&item)

	itemJSON, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(itemJSON))
}

// ItemView .
func ItemView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var item Item
	Db.First(&item, id)
	if item.ID == 0 {
		http.Error(w, "item record not found.", http.StatusNotFound)
		return
	}

	itemJSON, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(itemJSON))
}

// ItemDelete .
func ItemDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var item Item
	Db.First(&item, id)
	if item.ID == 0 {
		http.Error(w, "item record not found.", http.StatusNotFound)
		return
	}

	Db.Delete(&item, id)

	res, _ := json.Marshal(struct {
		Data string `json:"data"`
	}{"success"})

	fmt.Fprint(w, res)
}
