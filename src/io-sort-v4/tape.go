package main

import (
	"os"
)

type tape struct {
	// file io.ReadWriteSeeker
	file *os.File
}

// 从文件开头开始写入指定数据
func (t *tape) Write(p []byte) (b int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}
