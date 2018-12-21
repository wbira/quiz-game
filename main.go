package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
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

func main() {
	csvFileName := flag.String("csvFile", "problems.csv", "csvFile defines path to csv file with questions and answers for quiz")
	flag.Parse()
	lines := csvReader(csvFileName)
	problems := mapLinesToProblems(lines)

	points := 0
	for i, problem := range problems {
		fmt.Printf("Problem %v: %v\n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			points++
		}
	}

	fmt.Printf("You scored %v from %v points\n", points, len(problems))
}
