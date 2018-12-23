package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func csvReader(csvFileName *string) [][]string {
	file, error := os.Open(*csvFileName)
	if error != nil {
		exit(fmt.Sprintf("Failed to open files %s", *csvFileName))
	}

	csvReader := csv.NewReader(file)
	lines, error := csvReader.ReadAll()
	if error != nil {
		exit("Failed to parse a file")
	}
	return lines
}

func mapLinesToProblems(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func displayFinalMessage(points, problems int) {
	fmt.Printf("\nYou scored %v from %v points\n", points, problems)
}

func parseFlags() (*string, *int) {
	csvFileName := flag.String("csvFile", "problems.csv", "csvFile defines path to csv file with questions and answers for quiz")
	timeLimit := flag.Int("limit", 30, "defines time limit for single game")
	flag.Parse()
	return csvFileName, timeLimit
}

func main() {
	csvFileName, timeLimit := parseFlags()
	lines := csvReader(csvFileName)
	problems := mapLinesToProblems(lines)
	points := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, problem := range problems {
		fmt.Printf("Problem %v: %v\n", i+1, problem.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			displayFinalMessage(points, len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.answer {
				points++
			}
		}
	}
	displayFinalMessage(points, len(problems))
}
