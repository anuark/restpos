package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	db := GetDb()
	defer db.Close()

	// organization := Organization{Logo: "/static/fruit.png", Name: "Orange Zero"}
	// db.Create(&organization)
	// user := User{Username: "Ark", Email: "jaicof@gmail.com", OrganizationID: organization.ID}
	// db.Create(&user)

	// var user User
	// db.First(&user)
	// user.GenerateAuthKey()
	// user.Password = "asd123"
	// user.HashPassword()
	// db.Save(&user)

	if os.Getenv("REACT_APP_RESTPOS_HOST") == "" {
		log.Fatal("REACT_APP_RESTPOS_HOST env is not set.")
	}

	r := mux.NewRouter()
	r.Use(Cors, Auth, JSONContentType, Log)
	r.HandleFunc("/auth", UserAuth).Methods("POST", "OPTIONS")

	// order.go
	r.HandleFunc("/orders", OrderIndex).Methods("GET", "OPTIONS")
	r.HandleFunc("/orders", OrderCreate).Methods("POST", "OPTIONS")
	r.HandleFunc("/orders", OrderUpdate).Methods("PUT", "OPTIONS")
	r.HandleFunc("/orders", OrderDelete).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/orders/{id}", OrderView).Methods("GET", "OPTIONS")

	// item.go
	r.HandleFunc("/items", ItemIndex).Methods("GET", "OPTIONS")
	r.HandleFunc("/items", ItemCreate).Methods("POST", "OPTIONS")
	r.HandleFunc("/items", ItemUpdate).Methods("PUT", "OPTIONS")
	r.HandleFunc("/items", ItemDelete).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/items/{id}", ItemView).Methods("GET", "OPTIONS")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets/"))))
	http.ListenAndServe("localhost:81", r)
}
