package main

import (
	"encoding/csv"
	"fmt"
  "strconv"
	"log"
	"os"
)
func main() {
  // Open problems.csv
  f, err := os.Open("problems.csv")
  if err != nil {
    log.Fatal(err)
  }
  
  defer f.Close()
  // read the csv in utilizing "encoding/csv"
  csvReader := csv.NewReader(f)
  data, err := csvReader.ReadAll()
  if err != nil {
    log.Fatal(err)
  }
  nQuestions := len(data)
  answers := make([]int, nQuestions)
  
  

  fmt.Printf("Starting Quiz!\n")
  var answer int
  var question string
  
  for value := range data {
    question = data[value][0]
    fmt.Printf("Question %d: %s = ?\n", value+1, question)
    fmt.Scanf("%d\n", &answer)
    answers[value] = answer
  }

  var totalCorrect int = 0

  for value := range data {
    intVar, err := strconv.Atoi(data[value][1])
    if err != nil {
      log.Fatal(err)
    }

    if intVar == answers[value] {
      totalCorrect++
      fmt.Printf("Q%d: %s == %d :: Correct\n", value+1, data[value][0], intVar)
    } else {
      
      fmt.Printf("Q%d: %s != %d :: Incorrect\n", value+1, data[value][0], answers[value])
      fmt.Printf("     %s = %d\n\n", data[value][0], intVar)

    }
  }

  fmt.Printf("Score: %.1f%%\n", float32(totalCorrect) / float32(nQuestions) * 100.0)

}
