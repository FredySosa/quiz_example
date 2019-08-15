package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	// set flags
	csvPath = flag.String("csv", "problems.csv", "a csv file that contains the questions in the format: answer,question")
	limit   = flag.Int("limit", 30, "time limit for answer the quiz")
	shuffle = flag.Bool("shuffle", false, "flag that indicates if you want to shuffle the order of the questions in every compilation")
)

type (
	Problem struct {
		question string
		answer   string
	}
)

func NewProblem(question, answer string) *Problem {
	return &Problem{
		question: question,
		answer:   answer,
	}
}

func main() {
	flag.Parse()
	file, err := os.Open(*csvPath)
	if err != nil {
		log.Fatalf("cannot open the file [%s]. Err=%v", *csvPath, err)
	}
	reader := csv.NewReader(file)
	problems := make([]*Problem, 0)
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("cannot read some records from file [%s], err = %v", *csvPath, err)
		}
		problems = append(problems, NewProblem(row[0], row[1]))
	}

	if *shuffle {
		rand.Seed(time.Now().Unix())
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	goodAnswers := 0
	for index, problem := range problems {
		fmt.Print(fmt.Sprintf("Question #%d: %s. Answer: ", index+1, problem.question))
		answer := readEntry()
		if answer == problem.answer {
			goodAnswers++
		}
	}

	fmt.Println(fmt.Sprintf("Right answers %d of %d", goodAnswers, len(problems)))

}

func readEntry() string {
	var answer string
	_, err := fmt.Scan(&answer)
	if err != nil {
		return ""
	}
	return answer
}
