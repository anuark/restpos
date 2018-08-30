package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := GetDb()
	defer db.Close()

	organization := Organization{Logo: "/static/fruit.png", Name: "Orange Zero"}
	db.Create(&organization)
	user := User{Username: "Ark", Email: "jaicof@gmail.com", OrganizationID: organization.ID}
	db.Create(&user)

	r := mux.NewRouter()
	r.Use(Cors, Auth, JSONContentType)
	r.HandleFunc("/auth", UserAuth).Methods("POST", "OPTIONS")

	r.HandleFunc("/orders", OrderIndex).Methods("GET", "OPTIONS")
	r.HandleFunc("/orders", OrderCreate).Methods("POST", "OPTIONS")
	r.HandleFunc("/orders", OrderUpdate).Methods("PUT", "OPTIONS")
	r.HandleFunc("/orders", OrderDelete).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/orders/{id}", OrderView).Methods("GET", "OPTIONS")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets/"))))
	http.ListenAndServe("localhost:81", r)
}
