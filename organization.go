package main

import "github.com/jinzhu/gorm"

type Organization struct {
	gorm.Model
	Name, Logo string
}
