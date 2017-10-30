package main

import (
	_ "encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}

func usage() {
	fmt.Printf("usage: %s [inputfile]\n", os.Args[0])
	os.Exit(2)
}
