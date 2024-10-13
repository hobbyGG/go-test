package main

import (
	"log"
	"net/http"
	"os"

	// 导入本地包的方法。依据路径导入，起始路径是gomod所在位置，然后按文件路径访问即可
	// poker "http-server/src/command-line/v1"
	poker "github.com/hobbyGG/go-test/master/src/command-line/v1"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
