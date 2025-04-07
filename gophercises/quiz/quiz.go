package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
	"strings"
)

//type problem struct {
//q string
//a string
//}

func main() {

	filePath := flag.String("filename", "problem.csv", "path to csv file with question and answers")
	timelimit := flag.Int("limit", 10, "time after which quiz should exit")
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

	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
problemloop:
	for i := range lines {
		fmt.Printf("Problem: #%d: %s = \n", i+1, lines[i][0])
		answerch := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerch <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Time's up")
			break problemloop
		case answer := <-answerch:
			if strings.TrimSpace(lines[i][1]) == strings.TrimSpace(answer) {
				//fmt.Println("Correct answer")
				correct++
			}
		}
		//		if i == 2 {
		//			break
		//		}
	}
	fmt.Printf("Score %d out of %d \n", correct, len(lines))

}
