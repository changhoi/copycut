package main

import (
	"fmt"
	"io"
	"os"

	"github.com/changhoi/copycut/internal"
)

func main() {
	args := os.Args[1:]
	length := len(args)
	if length < 2 || length > 3 {
		fmt.Println("usage: copycut <src filename> <dst filename> [bytes]")
		os.Exit(1)
	}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}()

	var mib int64
	if length == 3 {
		mib = internal.MustAtoI(args[2])
	}

	src := internal.MustOpenFile(args[0])
	defer src.Close()

	dst := internal.MustCreateFile(args[1])
	defer dst.Close()

	if mib == 0 {
		info, err := src.Stat()
		if err != nil {
			panic(err)
		}

		mib = info.Size()
	}

	var bufSize int64 = 4096

	buf := make([]byte, bufSize)
	for mib > 0 {
		n, err := src.Read(buf[0:])
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		var content []byte
		if int64(n) > mib {
			content = buf[0:mib]
		} else {
			content = buf[0:n]
		}

		n, err = dst.Write(content)
		if err != nil {
			panic(err)
		}

		mib -= int64(n)
	}
	info, err := dst.Stat()
	if err != nil {
		panic(err)
	}

	fileSize := info.Size()

	fmt.Printf("Write file %d bytes\n", fileSize)
}
