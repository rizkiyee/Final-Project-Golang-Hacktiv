package main

import (
	"fga-final-project-mygram/config"
	"fga-final-project-mygram/router"
)

func main() {
	config.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
