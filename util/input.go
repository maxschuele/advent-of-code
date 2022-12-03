package util

import (
	"os"
	"path"
	"runtime"
	"strings"
)

func GetInput() string {
	_, filename, _, _ := runtime.Caller(1)
	pwd := path.Dir(filename)

	pathf, e := os.ReadFile(pwd + "/input.txt")
	if e != nil {
		panic(e)
	}

	return strings.TrimSpace(string(pathf))
}

func GetInputLines() []string {
	return strings.Split(GetInput(), "\n")
}
