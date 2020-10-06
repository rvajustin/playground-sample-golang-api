package main

import (
    "fmt"
    "log"
    "net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "ok")
    fmt.Println("Endpoint Hit: healthCheck")
}

func info(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "golang")
    fmt.Println("Endpoint Hit: info")
}

func handleRequests() {
    http.HandleFunc("/healthcheck", healthCheck)
    http.HandleFunc("/info", info)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}