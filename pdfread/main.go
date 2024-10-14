package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// lets declare a struct to save all contents as list
	type LoanSchedule struct {
		InstallNum    string
		DueDate       string
		Oprincipal    string
		InstallAmount string
		PrinsAmount   string
		IntAmount     string
		Cprincipal    string
	}
	// Open the PDF file
	filePath := "loanschedule.txt" // Replace with your PDF file path
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	lsList := []LoanSchedule{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//loop over each line and add to list
		ls := LoanSchedule{}
		for scanner.Scan() {
			line := scanner.Text()
			items := strings.Split(line, " ")
			ls.InstallNum = items[0]
			ls.DueDate = items[1]
			ls.Oprincipal = items[2]
			ls.InstallAmount = items[3]
			ls.PrinsAmount = items[4]
			ls.IntAmount = items[5]
			ls.Cprincipal = items[6]
			//fmt.Println(ls)
			lsList = append(lsList, ls)
		}
		//fmt.Println(lsList)
	}

	// calulate sum of intrest payable from 25th Installment to 60th Installment
	var sum float64
	for _, v := range lsList {
		installNum, _ := strconv.ParseInt(v.InstallNum,10, 64)
		if installNum >= 25 && installNum <= 60 {
			intAmount := strings.ReplaceAll(v.IntAmount, ",", "")
			i, err := strconv.ParseFloat(intAmount, 64)
			if err != nil {
				panic(err)
			}
			sum += i
		}
	}
	fmt.Println(sum)
    //702693 281521
	// Check for errors during the file reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
