package main

import "github.com/jinzhu/gorm"

// Item .
type Item struct {
	gorm.Model
	Name, ImagePath string
	Enabled         bool
}
