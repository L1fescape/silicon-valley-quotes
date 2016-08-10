package main

import (
  "fmt"
  "net/http"
  "math/rand"
)

var quotes = ReadFile("quotes.txt")

func Index(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "static/index.html")
}

func Quote(w http.ResponseWriter, r *http.Request) {
  quote := quotes[rand.Intn(len(quotes))]
  fmt.Fprintf(w, quote)
}
