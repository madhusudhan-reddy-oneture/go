package main

import (
	"fmt"
	"log"
	"os"

	"github.com/madhusudhan-reddy-oneture/gotbd/my-app/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	fmt.Println("Let's play Poker")
	fmt.Println("Type {Name} wins to record a win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
