package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	filePath := flag.String("filename", "random", "filepath obsolute path")
	flag.Parse()
	//fmt.Println(*filePath)
	f, err := os.Open(*filePath)

	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)

	lines, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	for i := range lines {
		fmt.Println("Problem: ", lines[i][0])
		fmt.Println("answer: ", lines[i][1])
	}
}
