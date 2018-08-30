package main

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Name, Email uint
}
