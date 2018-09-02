package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	userkey "restpos/pkg/userKey"
	"strconv"

	Context "github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type image struct {
	URL         string `json:"url"`
	Description string `json:"desc"`
}

// Item .
type Item struct {
	Model
	Name           string `json:"name"`
	ImagePath      string `json:"image_path"`
	Enabled        bool   `json:"enabled"`
	CategoryID     uint   `json:"category_id"`
	OrganizationID uint   `json:"organization_id"`
	Image          image  `gorm:"-" json:"image"`
}

// AfterFind .
func (i *Item) AfterFind() (err error) {
	if i.ImagePath[:4] == "http" { // it's already a url
		i.Image = image{i.ImagePath, i.Name}
	} else {
		url := os.Getenv("REACT_APP_RESTPOS_HOST") + i.ImagePath
		fmt.Println(url)
		// imageStruct := image{url, i.Name}
		i.Image = image{url, i.Name}
		// i.Image, err = json.Marshal(imageStruct)
	}
	return
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

	user := Context.Get(r, userkey.Key).(User)

	item.OrganizationID = user.OrganizationID
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
