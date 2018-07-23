package main

import (
	// "APIGateways/app/config"
	"APIGateways/app/routes"
	// "log"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	echo := route.Init()

	echo.Logger.Fatal(echo.Start(":" + port))

}
