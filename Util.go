package main

import (
  "bufio"
  "os"
  "log"
)

func ReadFile(fileName string) []string {
  lines := make([]string, 0, 50)

  file, err := os.Open(fileName)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  return lines
}

