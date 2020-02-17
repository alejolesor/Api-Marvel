package main

import (
	"fmt"

	"github.com/Api-Marvel/api/server"
	"github.com/Api-Marvel/internal/routes"
)

func main() {

	session := server.ConnectedDB()
	fmt.Println("conexion a db   "+"// ", session)
	router := routes.SetupRouter()
	router.Run()
}
