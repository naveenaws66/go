package main

import (
	"flag"
	"fmt"
)

func main() {

	filePath := flag.String("filename", "random", "filepath obsolute path")
	flag.Parse()
	fmt.Println(*filePath)

}