package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var csvFile string

func init() {
	flag.StringVar(&csvFile, "csv", "problems.csv", "path to csv file with problems")
	flag.Parse()
}

func main() {
	problems, err := os.ReadFile(csvFile)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(strings.NewReader(string(problems)))

	questionCount := 0
	score := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		questionCount++

		question, a := record[0], record[1]
		answer, err := strconv.Atoi(a)
		if err != nil {
			continue
		}

		var userInput int
		fmt.Print(question, " ")
		fmt.Scanf("%d", &userInput)
		if userInput == answer {
			score++
		}
	}

	fmt.Printf("%d/%d\n", score, questionCount)

}
