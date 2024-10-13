package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) Find(name string) *Player {

	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

// 对于io.Reader这个接口，可以想象它一个一个字节读取文件直到结束

func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}
