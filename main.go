package main

import (
	"encoding/csv"
	// "flag"
	"fmt"
	"log"
	"os"
)

func check(e error, msg string) {
    if e != nil {
        log.Fatal(e)
    }
}

func csvReader(filePath string) [][]string {
    file, err := os.Open(filePath)
    check(err, "Error opening file")

    reader := csv.NewReader(file)

    records, err := reader.ReadAll()
    check(err, "Error reading csv file")


    return records
}

func startQuiz(quiz [][]string) {
    var correctCount int
    for i, question := range quiz{
        var answer string
        fmt.Printf("Question %d: %s \n", (i + 1), question[0])
        fmt.Scan(&answer)
        if answer == question[1] {
            correctCount++
        }
    }
    fmt.Printf("You got %d questions correct", correctCount);
}

func main() {
    res := csvReader("./problems.csv")

    startQuiz(res)
}
