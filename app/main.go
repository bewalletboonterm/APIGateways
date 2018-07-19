package main

import (
	// "APIGateways/config"
	"APIGateways/app/routes"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	echo := route.Init()
	echo.Logger.Fatal(echo.Start(":" + port))

}
