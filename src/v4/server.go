package main

import (
	"fmt"
	"log"
	"net/http"
)

// 存储分数是一个接口,任何实现了取分数方法的类型都可以作为存储方式
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// 将函数重构为结构体
type PlayerServer struct {
	store PlayerStore // 存储分数
}

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {

}

// 函数实现改成结构体方法Handler
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}

}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	// map类型对于不存在的键对应的值是 key的默认值，int就是0，string就是空字符串""
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	fmt.Fprint(w, score)
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	err := http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
