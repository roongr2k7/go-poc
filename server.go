package main

import "net/http"
import "fmt"

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi Roong %s", r.URL.Path)
}
