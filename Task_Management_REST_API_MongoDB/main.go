package main

import (
	"TaskManagementRESTAPI_MongoDB/data"
	"TaskManagementRESTAPI_MongoDB/router"
)

func main() {
	r := router.RouterSetup()
	data.ConnectDB()
	r.Run("localhost:1000") // Start the server on localhost:1000

}
