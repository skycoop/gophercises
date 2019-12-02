package main

import (
	"bufio"
	"encoding/csv"
	"errors"
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
	timeLimit   time.Duration
}

func handleArgs() Options {
	problemFile := flag.String("problem-file", "quiz/data/problems.csv",
		"A CSV file containing the questions and answers")
	debug := flag.Bool("debug", false, "Whether to print debug logging")
	shuffle := flag.Bool("shuffle", false, "Whether to shuffle the problems loaded from the file")
	timeLimit := flag.Duration("time-limit", 30*time.Second,
		"Set the duration of the time limit for answering questions")
	flag.Parse()

	options := Options{
		problemFile: *problemFile,
		debug:       *debug,
		shuffle:     *shuffle,
		timeLimit:   *timeLimit,
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
	log.WithField("filename", filename).Debug("Loading problems from file")

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
		return nil, errors.New("no questions in file")
	}

	problems := make([]Problem, len(records))
	for _, record := range records {
		problem := Problem{
			question: record[0],
			answer:   strings.TrimSpace(record[1]),
		}
		problems = append(problems, problem)
	}

	if shuffle {
		log.Debug("Shuffling problem list")
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	log.WithField("count", len(problems)).Debug("Finished loading problems")
	return problems, nil
}

func collectUserInput(inputs chan string) {
	stdin := bufio.NewReader(os.Stdin)
	for {
		inputBytes, err := stdin.ReadBytes('\n')
		if err != nil {
			log.WithField("error", err).Fatal("Failed to parse user input")
		}
		input := strings.TrimSpace(string(inputBytes))
		log.WithField("input", input).Debug("Received user input")
		inputs <- input
	}
}

func main() {
	options := handleArgs()

	problems, err := loadProblems(options.problemFile, options.shuffle)
	if err != nil {
		log.WithField("error", err).Fatal("Failed to load problem file")
	}

	inputs := make(chan string, len(problems)+1)
	go collectUserInput(inputs)

	fmt.Printf("You will have %s to answer the questions. Hit enter to start\n", options.timeLimit)
	<-inputs
	timer := time.NewTimer(options.timeLimit)

	correct := 0

loop:
	for _, problem := range problems {
		fmt.Printf("%s? ", problem.question)
		select {
		case answer := <-inputs:
			if answer == problem.answer {
				correct++
			}
		case <-timer.C:
			fmt.Print("Time's up!\n")
			break loop
		}
	}

	fmt.Printf("You answered %d correctly out of %d.\n", correct, len(problems))
}
