package main

import (
	"io"
	"testing"
)

// 一般不允许测试私有的东西，这会导致紧耦合，阻碍重构
func TestTape(t *testing.T) {
	file, clean := createTempFile(t, "123456")
	defer clean()

	tape := &tape{file: file}

	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
