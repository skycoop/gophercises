package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"strings"
	"time"
)

func init() {
	// TODO: Figure out why the time sometimes gets stuck on a constant value resulting in completely deterministic shuffling
	rand.Seed(time.Now().UnixNano())
}

type Options struct {
	problemFile string
	debug       bool
	shuffle     bool
}

func handleArgs() Options {
	problemFile := flag.String("problem-file", "quiz/data/problems.csv", "A CSV file containing the questions and answers")
	debug := flag.Bool("debug", false, "Whether to print debug logging")
	shuffle := flag.Bool("shuffle", false, "Whether to shuffle the problems loaded from the file")
	flag.Parse()

	options := Options{
		problemFile: *problemFile,
		debug:       *debug,
		shuffle:     *shuffle,
	}

	if options.debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	return options
}

type Problem struct {
	question string
	answer   string
}

func loadProblems(filename string, shuffle bool) ([]Problem, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, fmt.Errorf("no questions in file")
	}

	var problems []Problem
	for _, record := range records {
		problem := Problem{
			question: record[0],
			answer:   strings.TrimSpace(record[1]),
		}
		problems = append(problems, problem)
	}

	if shuffle {
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	return problems, nil
}

func main() {
	options := handleArgs()

	problems, err := loadProblems(options.problemFile, options.shuffle)
	if err != nil {
		log.Fatalf("Failed to load problem file: %s", err)
	}

	correct := 0
	for _, problem := range problems {
		log.WithFields(log.Fields{
			"question": problem.question,
			"answer":   problem.answer,
		}).Debug("Asking problem")

		fmt.Printf("%s? ", problem.question)
		var answer string
		_, err := fmt.Scanf("%s\n", &answer)
		if err != nil {
			log.Fatalf("Failed to parse user input: %s", err)
		}

		if answer == problem.answer {
			correct++
		}
	}
	fmt.Printf("You answered %d correctly out of %d.\n", correct, len(problems))
}
