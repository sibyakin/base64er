package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	filenames := os.Args[1:]
	for _, filename := range filenames {
		in, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("Unable to open %s with error %s\n", filename, err)
			continue
		}

		out, err := os.Create(filename + ".json")
		if err != nil {
			fmt.Printf("Unable to create %s with error %s\n", filename+".json", err)
			continue
		}
		defer out.Close()

		buf := new(bytes.Buffer)

		writer := base64.NewEncoder(base64.StdEncoding, buf)
		writer.Write(in)
		defer writer.Close()

		result := "{ \"name\": " + filename + ", \"image64\": " + buf.String() + "\" }"

		_, err = io.Copy(out, strings.NewReader(result))
		if err != nil {
			fmt.Printf("Unable to convert: %s!\n", err)
		}
	}
}

func usage() {
	fmt.Printf("usage: %s [inputfile]\n", os.Args[0])
	os.Exit(2)
}
