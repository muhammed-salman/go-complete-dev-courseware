package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	} else {
		io.Copy(os.Stdout, file)
	}
}
