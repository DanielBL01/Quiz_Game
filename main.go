package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFile := flag.String("problems", "problems.csv", "Answer simple math questions from csv file")
	flag.Parse()

	// Dereferencing pointer calls CSV file name
	file, err := os.Open(*csvFile)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s\n", *csvFile)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	// records is now a 2D slices or slice of slices [][]string

	if err != nil {
		fmt.Println("Failed")
		os.Exit(1)
	}

	problems := parseRecord(records)

	// explicit value init, go will know it's int type
	numCorrect := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)

		// explicit type init, empty string
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			fmt.Println("Correct!")
			numCorrect++
		} else if answer != problem.answer {
			fmt.Println("Incorrect!")
		}
	}

	fmt.Printf("\nYou got %d questions out of %d questions!", numCorrect, len(problems))

}

func parseRecord(records [][]string) []problem {
	result := make([]problem, len(records))
	// Similar to enum in Python
	for i, record := range records {
		result[i] = problem{
			question: record[0], //First Column question
			answer:   record[1], //Second Column answer
		}
	}
	return result
}

type problem struct {
	question string
	answer   string
}
