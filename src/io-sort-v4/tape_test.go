package main

import (
	"io"
	"testing"
)

// tape直接是一个struct不是一个接口，我们一般不允许直接测试struct，应该实现一个interface并传入struct
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
