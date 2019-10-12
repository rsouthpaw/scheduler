package main

import (
	"./api"
	"./base"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	serverDetails, err := base.SetupServer(base.SERVER_TYPE_LOCALHOST)
	if err != nil {
		log.Println("ERROR SETTING UP SERVER")
		return
	}
	log.Println(serverDetails)
	http.ListenAndServe(base.PORT, api.GetRouter())
}
