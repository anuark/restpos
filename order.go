package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

// Order .
type Order struct {
	gorm.Model
	CustomerID uint
}

// OrderCreate .
func OrderCreate(w http.ResponseWriter, r *http.Request) {

}

// OrderIndex .
func OrderIndex(w http.ResponseWriter, r *http.Request) {

}

// OrderUpdate .
func OrderUpdate(w http.ResponseWriter, r *http.Request) {

}

// OrderView .
func OrderView(w http.ResponseWriter, r *http.Request) {
}

// OrderDelete .
func OrderDelete(w http.ResponseWriter, r *http.Request) {

}
