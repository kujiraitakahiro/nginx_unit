package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Test G0 for nginx unit")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":7777", nil)
  //nginx unitでは http.ListenAndServe を unit.ListenAndServe に変更する
}
