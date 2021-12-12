package internal

import (
	"os"
	"strconv"
)

func MustOpenFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return f
}

func MustCreateFile(filename string) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	return f
}

func MustAtoI(a string) int64 {
	ret, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		panic(err)
	}

	return ret
}
