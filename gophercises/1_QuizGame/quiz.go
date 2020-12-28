// Gophercise 1 - Quiz Game - https://gophercises.com/exercises/quiz
//
// Packages used
//		log  -- logging -- https://www.datadoghq.com/blog/go-logging/
//  Logging -- zerolog -- https://github.com/rs/zerolog   ($ go get github.com/rs/zerolog)
//  Testing --
//		flag -- command-line args -- https://blog.alexellis.io/5-keys-to-a-killer-go-cli/
//      pflag -- replacement for Go's "flag" package (supports POSIX style flags)
//               https://godoc.org/github.com/spf13/pflag#BoolVarP
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	flag "github.com/spf13/pflag" //  replacement for Go's flag package, implementing POSIX style flags
	zlog "github.com/rs/zerolog/log"    // level-based logger based on zerolog; $ go get gopkgs.sas.com/zlog
)

var logger = zlog.L()

type questionAnswer struct {
	question string
	answer   string
}

func main() {
	logger.Info("Gophercise 1 - Quiz Game - https://gophercises.com/exercises/quiz")

	// process command line args
	var showHelp bool
	flag.BoolVarP(&showHelp, "help", "h", false, "Display usage help")
	var csvFilename string
	flag.StringVarP(&csvFilename, "file", "f", "problems.csv", "Comma-separated file of question-answer pairs")
	var verbose bool
	flag.BoolVarP(&verbose, "verbose", "v", false, "Verbose mode; log actions in great detail")
	var version1 bool
	flag.BoolVarP(&version1, "version1", "1", false, "Run version one of the quiz app (no timer)")

	flag.Parse()
	if showHelp {
		//NOTE: All this might be overkill because pflag can generate the bulk it (VERIFY THIS)
		logger.Fatal("Gophercise 1 - Quiz Game\n\n" +
			"Usage:\n" +
			"\t$ quiz [--help] [--file csvFile]\n" +
			"where,\n" +
			"\t--help       Display this text (shorthand: -h)\n" +
			"\t--file       CSV file to use as input")
	}

	// load CSV file
	f, err := os.Open(csvFilename)

	defer f.Close()
	if err != nil {
		logger.Fatal(err)
	}
	csvReader := csv.NewReader(bufio.NewReader(f))
	var questionsAndAnswers []questionAnswer
	var lineNo int
	for {
		lineNo++
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Fatal("line", lineNo, ":", err)
		}
		if len(row) != 2 {
			logger.Fatal("line", lineNo, ": Expecting two values separated by a comma, but read:", row)
		}
		logger.Debug("line", lineNo, ":", row)
		}
		questionsAndAnswers = append(questionsAndAnswers, questionAnswer{row[0], row[1]})
	}
	var questionCount := len(questionsAndAnswers)
	if verbose {
		logger.Println("Read", questionCount, "questions from", csvFilename, ". Contents of questionsAndAnswers:\n", questionsAndAnswers)
	}

	// shuffle the questions
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(questionCount, func(i, j int) {
		questionsAndAnswers[i], questionsAndAnswers[j] = questionsAndAnswers[j], questionsAndAnswers[i]
	})

	var correctAnswers int

	if version1 {
		fmt.Printf("\nGophercise 1 - Quiz Game - Version 1 (no timer)\n\n")

		// Query the user
		var response string
		for i, qa := range questionsAndAnswers {
			fmt.Printf("(%d) What is %s? ", i, qa.question)
			fmt.Scanf("%s\n", &response)
			if strings.EqualFold(qa.answer, response) { // case-insensitive string compare
				fmt.Println("\t\tRight!")
				correctAnswers++
			} else {
				fmt.Println("\t\t\t\tBZZZZZZZZZZZ!")
			}
		}
	} else {
		maxTime := questionCount * 5 // 5 seconds per question
		fmt.Printf("\nGophercise 1 - Quiz Game - Version 2\nYou will have %v seconds to answer %v qusetions... GO!!\n"+
			"Press ENTER to begin.\n\n", maxTime, questionCount)

	}

	fmt.Printf("\n========= End of Quiz =========\n"+
		"You answered %d out of %d correctly.\n"+
		"Score: %d", correctAnswers, questionCount, 100*correctAnswers/questionCount)
}
