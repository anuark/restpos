package main

import "github.com/jinzhu/gorm"

// OrderItem .
type OrderItem struct {
	gorm.Model
	OrderID uint
	ItemID  uint
}
