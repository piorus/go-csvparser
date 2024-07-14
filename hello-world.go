package main

import (
	"bufio"
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

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
  check(scanner.Err())
}

