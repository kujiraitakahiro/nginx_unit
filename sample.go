package main

import (
  "fmt"
  "net/http"
  //nginx unitで使う場合、"nginx/unit" もimportするべし
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Test Go for nginx unit")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":7777", nil)
  //nginx unitでは http.ListenAndServe を unit.ListenAndServe に変更する
}
