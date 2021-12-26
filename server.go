package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net"
	"net/http"
	"os"
	"rest-api/config"
	"rest-api/route"
)

func main() {

	config.LoadLocalEnv()

	address := net.JoinHostPort(os.Getenv("ADDRESS"), os.Getenv("PORT"))
	fmt.Printf("Server is running at %v \n", address)
	err := http.ListenAndServe(address, route.ConfigureRoutes())
	if err != nil {
		return
	}

}
