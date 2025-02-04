package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("quiz.csv")

	if err != nil {
		fmt.Errorf("Error while loading quiz file....", err)
		return
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Errorf("Error while reading file....", err)
		return
	}

	for i, record := range records {
		fmt.Printf("\nProblem #%d: %s + %s = ", i+1, record[0], record[1])

		var answer string
		fmt.Scanln(&answer)

		if answer != record[2] {
			fmt.Printf("\nYou scored %d out of %d", i+1, len(records))
			return
		}
	}

}
