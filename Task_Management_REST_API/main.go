package main

import (
	"Task_Management_REST_API/router"
)

func main() {
	r := router.RouterSetup()
	r.Run("localhost:1000") // Start the server on port 8080

}
