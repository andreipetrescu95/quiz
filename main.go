package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
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

	flag.Parse()

	data, err := ioutil.ReadFile(*fileName)
	inputReader := bufio.NewReader(os.Stdin)

	noQuestions := 0

	for i := 0; i < len(data); i++ {
		if data[i] == 10 {
			noQuestions++
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(data)))

	correctAnswers := 0

	fmt.Print("Press enter to start the game!")
	inputReader.ReadString('\n')

	go timer(*limit, &correctAnswers, noQuestions)

	i := 0

	for {
		record, err := r.Read()

		if err == io.EOF {
			fmt.Println("You scored", correctAnswers, "out of", noQuestions)
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("Problem #", i, ": ", record[0], " = ")
		answer, err := inputReader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		i++

		if strings.TrimSpace(answer) == record[1] {
			correctAnswers++
		}
	}

}
