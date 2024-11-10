package main

import (
    "fmt"
    "net/http"
    "os"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    hostname, err := os.Hostname()
    fmt.Fprintf(w, "Hello World ")
    if err == nil {
      fmt.Fprintf(w, hostname)
   }
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":3000", nil)
}
