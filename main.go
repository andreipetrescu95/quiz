package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

func timer(limit int, correctAnswers *int, noQuestions int) {
	start := time.Now()

	for {
		if (time.Now().Unix() - start.Unix()) > int64(limit) {
			fmt.Print("\n")
			fmt.Println("You scored", *correctAnswers, "out of", noQuestions)
			os.Exit(0)
		}
	}
}

func main() {
	fileName := flag.String("filename", "problems.csv", "Filename for the quiz with questions and answers")
	limit := flag.Int("limit", 30, "Time limit for the quiz")
	shuffle := flag.Bool("shuffle", false, "pass true to shuffle the quiz, and false, to remain as in the file")

	flag.Parse()

	data, err := ioutil.ReadFile(*fileName)
	inputReader := bufio.NewReader(os.Stdin)

	if err != nil {
		log.Fatal(err)
	}

	questions := strings.Split(string(data), "\n")

	if questions[len(questions)-1] == "" {
		questions = questions[0 : len(questions)-1]
	}

	correctAnswers := 0

	fmt.Print("Press enter to start the game!")
	inputReader.ReadString('\n')

	if *shuffle {
		rand.Seed(time.Now().UTC().UnixNano())

		sort.Slice(questions, func(i, j int) bool {
			randNumber := rand.Int() % 2
			randBool := false

			if randNumber == 1 {
				randBool = true
			}

			return randBool
		})
	}

	go timer(*limit, &correctAnswers, len(questions))

	for i := 0; i < len(questions); i++ {
		r := csv.NewReader(strings.NewReader(questions[i]))
		record, err := r.Read()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("Problem #", i, ": ", record[0], " = ")
		answer, err := inputReader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(strings.TrimSpace(answer)) == strings.ToLower(strings.TrimSpace(record[1])) {
			correctAnswers++
		}
	}

	fmt.Println("You scored", correctAnswers, "out of", len(questions))
}
