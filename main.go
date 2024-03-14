package main

import (
	"log"
	"net/http"
	"platform_x/router"
)

func main() {
	ginRouter := router.InitObjects()
	err := http.ListenAndServe(":8080", ginRouter)
	if err != nil {
		log.Fatal("Unable to start sever, error: ", err)
	}
	log.Println("Server started at port 8080")
}
