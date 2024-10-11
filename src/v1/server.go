package main

import (
	"fmt"
	"log"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {

	// r.URL.Path 返回请求的路径
	player := r.URL.Path[len("/palyers/"):]

	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {

	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}

func main() {
	handler := http.HandlerFunc(PlayerServer)
	err := http.ListenAndServe(":5000", handler)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
