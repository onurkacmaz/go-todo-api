package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"rest-api/config"
	"rest-api/database"
	"rest-api/route"
)

func main() {

	config.LoadLocalEnv()

	database.Migrate()

	address := net.JoinHostPort(os.Getenv("ADDRESS"), os.Getenv("PORT"))
	fmt.Printf("Server is running at %v \n", address)
	err := http.ListenAndServe(address, route.RegisterRoutes())
	if err != nil {
		return
	}

}
