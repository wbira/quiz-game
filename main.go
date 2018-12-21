package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func main() {
	csvFileName := flag.String("csvFile", "problems.csv", "csvFile defines path to csv file with questions and answers for quiz")
	flag.Parse()

	file, error := os.Open(*csvFileName)
	if error != nil {
		exit(fmt.Sprintf("Failed to open files %s", *csvFileName))
	}

	csvReader := csv.NewReader(file)
	lines, error := csvReader.ReadAll()
	if error != nil {
		exit("Failed to parse a file")
	}

	fmt.Println(lines)
}
