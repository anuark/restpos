package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Item .
type Item struct {
	Model
	Name       string `json:"name"`
	ImagePath  string `json:"image_path"`
	Enabled    bool   `json:"enabled"`
	CategoryID uint   `json:"category_id"`
}

// ItemIndex .
func ItemIndex(w http.ResponseWriter, r *http.Request) {
	var items []Item
	Db.Find(&items)

	itemsJSON, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err)
	}

	var count int
	Db.Table("item").Count(&count)
	w.Header().Add("X-Total-Count", strconv.Itoa(count))

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
