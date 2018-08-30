package main

import "github.com/jinzhu/gorm"

// Category .
type Category struct {
	gorm.Model
	Name string
}
