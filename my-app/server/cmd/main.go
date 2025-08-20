package main

import (
	"log"
	"net/http"

	"github.com/madhusudhan-reddy-oneture/gotbd/my-app/server"
)

func main() {
	store := NewInMemoryPlayerStore()
	server := server.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
