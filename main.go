package main

import (
	//bencode "github.com/jackpal/bencode-go"

	"fmt"
	"io"
	"os"
)

func main() {
	in_file_path := os.Args[1]

	file, err := os.Open(in_file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
    
    // Import torrent and parse file

}
