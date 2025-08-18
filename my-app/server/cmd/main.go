package main

import (
	"log"
	"net/http"

	"github.com/madhusudhan-reddy-oneture/gotbd/my-app/server"
)

type InMemoryPlayerStore struct{}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 0
}

func (s *InMemoryPlayerStore) RecordWin(name string) {

}

func main() {
	server := &server.PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
