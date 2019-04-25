package main

import (
	"fmt"
	"log"
	"net/http"
	"smilo-status/routers"
)

func main() {
	router := routers.NewRouter()
	fmt.Printf("Listening on localhost:5000\n")
	log.Fatal(http.ListenAndServe(":5000", router))
}
