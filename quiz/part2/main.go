package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var csvFile string
var timeLimit int

func init() {
	flag.StringVar(&csvFile, "csv", "problems.csv", "path to csv file with problems")
	flag.IntVar(&timeLimit, "limit", 10, "time limit for size in seconds")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
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

func shuffle(problems []problem) []problem {
	for i := range problems {
		swapI := rand.Intn(len(problems))
		problems[i], problems[swapI] = problems[swapI], problems[i]
	}
	return problems
}

func main() {
	problems := shuffle(generateProblemsArray())
	t := time.NewTimer(time.Duration(timeLimit) * time.Second)
	answerCh := make(chan int)
	correct := 0

problemLoop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ? ", i+1, problem.question)
		go func() {
			var a int
			// scanf is blocking code and so we don't want it blocking the main thread
			// (if it did then even if the timer expires we'd have to wait for scanf to stop blocking before ending the program)
			fmt.Scanf("%d", &a)
			answerCh <- a
		}()

		select {
		case <-t.C:
			// timer expires (break out of the for loop)
			break problemLoop
		case answer := <-answerCh:
			// received an answer (before timer expires)
			if answer == problem.answer {
				correct++
			}
		}
	}

	fmt.Printf("Score: %d/%d", correct, len(problems))
}
