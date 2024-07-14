package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  file, err := os.Open("./customers-1000.csv")
  check(err)
  defer file.Close()

  csvReader := csv.NewReader(file)
  records, err := csvReader.ReadAll()
  check(err)

  for _, row := range records {
    fmt.Println(row[3])
  }
}

