package main

import (
	"log"
	"net/http"

	"github.com/madhusudhan-reddy-oneture/gotbd/my-app/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5050", server); err != nil {
		log.Fatalf("could not listen on port 5050, %v", err)
	}
}
