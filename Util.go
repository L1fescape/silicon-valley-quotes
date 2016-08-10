package main

import (
  "bufio"
  "os"
  "log"
)

func ReadFile(fileName string) []string {
  lines := make([]string, 83) // todo don't hard-code the file length

  file, err := os.Open(fileName)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  counter := 0
  for scanner.Scan() {
    lines[counter] = scanner.Text()
    counter++
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  return lines
}

