package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
    fmt.Fprintf(w, "%s", hostname)
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}
