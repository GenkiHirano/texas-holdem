package main

import (
	"log"
	"net/http"

	poker "github.com/GenkiHirano/tdd-go.git"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server, err := poker.NewPlayerServer(store)

	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
