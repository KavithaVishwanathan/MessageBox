package main

import (
	"log"
    "net/http"
    "fmt"
    _ "github.com/KavithaVishwanathan/messagebox/db"
)

func homePage(rw http.ResponseWriter, request *http.Request) {
    rw.Write([]byte("Hello world."))
    fmt.Fp
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8081", nil))
}


func main() {
    handleRequests()
}