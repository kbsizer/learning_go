// Gophercise 1 - Quiz Game - https://gophercises.com/exercises/quiz
//
// Packages used
//		log  -- logging -- https://www.datadoghq.com/blog/go-logging/
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

	"github.com/spf13/pflag" //  replacement for Go's flag package, implementing POSIX style flags
)

type questionAnswer struct {
	question string
	answer   string
}

func main() {
	log.Print("Gophercise 1 - Quiz Game - https://gophercises.com/exercises/quiz")

	// process command line args
	var showHelp bool
	pflag.BoolVarP(&showHelp, "help", "h", false, "Display usage help")
	var csvFilename string
	pflag.StringVarP(&csvFilename, "file", "f", "problems.csv", "Comma-separated file of question-answer pairs")
	var verbose bool
	pflag.BoolVarP(&verbose, "verbose", "v", false, "Verbose mode; log actions in great detail")
	var version1 bool
	pflag.BoolVarP(&version1, "version1", "1", false, "Run version one of the quiz app (no timer)")

	pflag.Parse()
	if showHelp {
		//NOTE: All this might be overkill because pflag can generate the bulk it (VERIFY THIS)
		log.Fatal("Gophercise 1 - Quiz Game\n\n" +
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
		log.Fatal(err)
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
			log.Fatal("line", lineNo, ":", err)
		}
		if len(row) != 2 {
			log.Fatal("line", lineNo, ": Expecting two values separated by a comma, but read:", row)
		}
		if verbose {
			log.Println("line", lineNo, ":", row)
		}
		questionsAndAnswers = append(questionsAndAnswers, questionAnswer{row[0], row[1]})
	}
	questionCount := len(questionsAndAnswers)
	if verbose {
		log.Println("Read", questionCount, "questions from", csvFilename, ". Contents of questionsAndAnswers:\n", questionsAndAnswers)
	}

	// shuffle the questions
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(questionCount, func(i, j int) {
		questionsAndAnswers[i], questionsAndAnswers[j] = questionsAndAnswers[j], questionsAndAnswers[i]
	})

	if version1 {
		fmt.Println("\nGophercise 1 - Quiz Game - Version 1 (no timer)\n")

		// Query the user
		var response string
		for i, qa := range questionsAndAnswers {
			fmt.Printf("(%d) What is %s? ", i, qa.question)
			fmt.Scanf("%s\n", &response)
			if strings.EqualFold(qa.answer, response) { // case-insensitive string compare
				fmt.Println("   Right!")
			} else {
				fmt.Println("   BZZZZZZZZZZZ!")
			}
		}
	} else {
		maxTime := questionCount * 5 // 5 seconds per question
		fmt.Println("\nGophercise 1 - Quiz Game - Version 2\nYOU HAVE", maxTime, "SECONDS... GO!!\n")
	}
}
