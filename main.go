package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	filenames := os.Args[1:]
	for _, filename := range filenames {
		in, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("Unable to open %s with error %s", filename, err)
			continue
		}

		out, err := os.Create(filename + ".based64.json")
		if err != nil {
			fmt.Printf("Unable to create %s with error %s", filename+".based64.json", err)
			continue
		}
		defer out.Close()

		buf := new(bytes.Buffer)

		writer := base64.NewEncoder(base64.StdEncoding, buf)
		writer.Write(in)
		defer writer.Close()

		result := "{ \"name\": " + filename + ", \"image64\": " + buf.String() + "\" }"

		fmt.Printf(result + "\n")
	}
}

func usage() {
	fmt.Printf("usage: %s [inputfile]\n", os.Args[0])
	os.Exit(2)
}
