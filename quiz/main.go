package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile := flag.String("csv", "quiz.csv", "CSV file in the format of question,answer")
	timeLimit := flag.Int("limit", 10, "The time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFile)

	if err != nil {
		Exit(fmt.Sprintf("Error while loading quiz, file:%s, error:%s", *csvFile, err))
		os.Exit(1)
	}

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		Exit("Error while parsing quiz file")
	}

	quiz := ParseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correctAnwser := 0

	for i, q := range quiz {
		fmt.Printf("Problem #%d: %s = \n", i, q.question)
		answerCh := make(chan string)
		go func() {
			var userAnswer string
			fmt.Scanf("%s\n", &userAnswer)
			answerCh <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correctAnwser, len(quiz))
			return
		case userAnswer := <-answerCh:
			if userAnswer == q.answer {
				correctAnwser++
			}

		}
	}
	fmt.Printf("You scored %d out of %d.\n", correctAnwser, len(quiz))
}

func ParseLines(lines [][]string) []quiz {
	ret := make([]quiz, len(lines))
	for i, line := range lines {
		ret[i] = quiz{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type quiz struct {
	question string
	answer   string
}

func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
