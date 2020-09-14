package main

import (
        "net/http"
        "fmt"
        "log"
)

func main() {
        http.HandleFunc("/", hello)
        http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r  *http.Request) {
        w.Write([]byte("Hello, World!!"))
}
