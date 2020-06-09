package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// Initialize default settings for our file and timer limit
	csvFile := flag.String("problems", "problems.csv", "Answer simple math questions from csv file")
	timeLimit := flag.Int("Limit", 30, "The time limit for the quiz in seconds")
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

	// convert an integer number of units to a Duration. Init timer value from user
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// explicit value init, go will know it's int type
	numCorrect := 0
	for i, problem := range problems {
		select {
		case <-timer.C:
			fmt.Printf("\nYou got %d out of %d questions correct!", numCorrect, len(problems))
			return // using the break statement here would only break out of select
		default:
			fmt.Printf("Problem #%d: %s = ", i+1, problem.question)

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

	}

	fmt.Printf("\nYou got %d out of %d questions correct!", numCorrect, len(problems))

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
