package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// 对应了json格式
type Player struct {
	Name string
	Wins int
}

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(file *os.File) *FileSystemPlayerStore {
	file.Seek(0, 0)
	league, _ := NewLeague(file)

	// 将database替换成tape，tape只实现了write功能，所以主结构体将改为Writer，但是tape内部为RWS，支持先前的所有操作
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() League {

	// 每次都把文件指针放到最前
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {

	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	// 重置文件指针位置并写回
	f.database.Encode(f.league)
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordinWin(name string)
	GetLeague() League
}

// 我们更改了 PlayerServer 的第二个属性，删除了命名属性 router http.ServeMux，并用 http.Handler 替换了它；这被称为 嵌入。
type PlayerServer struct {
	store        FileSystemPlayerStore
	http.Handler // 直接把某个类型作为结构体类型相当于子类，我们可以调用嵌入类型的方法
}

// Go 没有提供典型的，类型驱动的子类化概念，但它具有通过在结构或接口中嵌入类型来“借用”一部分实现的能力。

// 新业务代码
func NewPlayerServer(store FileSystemPlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

// 前业务代码
// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	p.router.ServeHTTP(w, r)
// }

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}
