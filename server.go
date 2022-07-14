package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, getPlayerScore(player))
}

func getPlayerScore(n string) string {
	if n == "Pepper" {
		return "20"
	}

	if n == "Floyd" {
		return "10"
	}

	return ""
}
