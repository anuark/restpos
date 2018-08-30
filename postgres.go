package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Db .
var Db *gorm.DB
var once sync.Once

// GetDb .
func GetDb() *gorm.DB {
	once.Do(func() {
		mig := flag.Bool("migrate", false, "Run migrations")
		flag.Parse()
		verbose := flag.Bool("v", false, "Verbose mode")
		var err error
		Db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 dbname=restpos user=postgres password=asd123 sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		Db.LogMode(*verbose)
		Db.SingularTable(true)

		if *mig {
			migrate()
			fmt.Println("Migrations runned")
			os.Exit(0)
		}
	})

	return Db
}

// Migrate .
func migrate() {
	Db.AutoMigrate(&User{}, &Organization{}, &Item{}, &Category{}, &Order{}, &Customer{}, &OrderItem{})

	Db.Model(&User{}).AddForeignKey("organization_id", "organization(id)", "RESTRICT", "CASCADE")
	Db.Model(&Item{}).AddForeignKey("category_id", "category(id)", "CASCADE", "CASCADE")
	Db.Model(&Order{}).AddForeignKey("customer_id", "customer(id)", "RESTRICT", "CASCADE")
	Db.Model(&OrderItem{}).AddForeignKey("order_id", "order(id)", "CASCADE", "CASCADE")
}
