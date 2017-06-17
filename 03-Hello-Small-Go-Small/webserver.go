package main

import (
	"math/rand"
    "fmt"
    "net/http"
	"time"
)

var serverID = rand.New(
	rand.NewSource(
		time.Now().UnixNano())).Intn(999999)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "var id = ", serverID, ";", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}