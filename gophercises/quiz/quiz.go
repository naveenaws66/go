package main

import (
	"encoding/csv"
	"flag"
	"os"
	"fmt"
)

//type problem struct {
	//q string
	//a string
//}

func main() {

	filePath := flag.String("filename", "problem.csv", "path to csv file with question and answers")
	flag.Parse()

	f, err := os.Open(*filePath)

	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)

	lines, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	correct := 0
	
	for i := range lines {
	 	fmt.Printf("Problem: #%d: %s = \n", i+1, lines[i][0])
		var answer string
		fmt.Scanf("%s\n", &answer)
		if lines[i][1] == answer {
			//fmt.Println("Correct answer")
			correct++
		}

//		if i == 2 {
//			break
//		}
	 }
	 fmt.Printf("Score %d out of %d \n", correct, len(lines))

}
