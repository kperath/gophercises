package main

import (
	"encoding/csv"
	"flag"
	"fmt"
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

type problem struct {
	question string
	answer   int
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generateProblemsArray() []problem {
	file, err := os.ReadFile(csvFile)
	logError(err)
	r := csv.NewReader(strings.NewReader(string(file)))

	lines, err := r.ReadAll() // array of lines
	logError(err)

	problems := make([]problem, len(lines))

	for i, line := range lines {
		answerNum, err := strconv.Atoi(line[1])
		logError(err)

		problems[i] = problem{
			question: line[0],
			answer:   answerNum,
		}
	}

	return problems
}

func main() {
	problems := generateProblemsArray()
	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ? ", i+1, problem.question)
		var answer int
		fmt.Scanf("%d", &answer)
		if answer == problem.answer {
			correct++
		}
	}

	fmt.Printf("Score: %d/%d", correct, len(problems))
}
