package main

import (
	"log"

	"github.com/Api-Marvel/internal/services"

	"github.com/Api-Marvel/internal/databases"
	"github.com/Api-Marvel/internal/routes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//gorm.Open("mysql", "lagash:.a@vMztxya.hS6b@tcp(db4free.net:3306)/marketplacelag?charset=utf8&parseTime=True")
	db, err := databases.CreateDatabase()
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	services.Db(db)
	router := routes.SetupRouter()
	router.Run()
}
