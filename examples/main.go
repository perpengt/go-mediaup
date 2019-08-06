package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/perpengt/mediaup/pkg/mediaup"

	"github.com/perpengt/ids"
)

func readConsole () (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Token: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
}

func main() {
	token, err := readConsole()
	if err != nil {
		panic(err)
	}

	fp, err := os.OpenFile("./sample.png", os.O_RDONLY, 0600)
	if err != nil {
		fmt.Printf("Failed to load sample image: %s\n", err.Error())
		os.Exit(1)
	}
	defer fp.Close()

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		fmt.Printf("Failed to read the sample image!\n")
		os.Exit(1)
	}

	idb, err := mediaup.UploadImageBytes("http://localhost:8075/", token, data)
	if err != nil {
		fmt.Printf("Upload failed: %s\n", err)
		os.Exit(1)
	}

	id := &ids.ID{}

	err = id.Scan(idb)
	if err != nil {
		fmt.Printf("Invalid ID: %s\n", err)
	}

	fmt.Printf("Uploaded successfully! See http://localhost:8075/%s@1000w.png\n", id.URIString())
}
