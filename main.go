package main

import (
	"didux-status/managers"
	"didux-status/routers"
	"fmt"
	"gopkg.in/robfig/cron.v2"
	"log"
	"net/http"
)

func cronJobs() {
	s := cron.New()

	s.AddFunc("@every 20s", func() {
		managers.GetPublicBlocks()
	})

	s.Start()
}

func main() {
	router := routers.NewRouter()
	fmt.Printf("Listening on localhost:5000\n")
	cronJobs()
	//fmt.Scanln()
	log.Fatal(http.ListenAndServe(":5000", router))
}
