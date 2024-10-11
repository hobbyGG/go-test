package main

import (
	"fmt"
	"log"
	"net/http"
)

// 存储分数是一个接口,任何实现了取分数方法的类型都可以作为存储方式
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// 将函数重构为结构体
type PlayerServer struct {
	store PlayerStore // 存储分数
}

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

// 函数实现改成结构体方法
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	w.WriteHeader(http.StatusNotFound)

	fmt.Fprint(w, p.store.GetPlayerScore(player))
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
	server := &PlayerServer{&InMemoryPlayerStore{}}

	err := http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
