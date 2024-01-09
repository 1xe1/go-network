package main

import (
	"fmt"
	"net/http"
)

func main() {
    fs := http.FileServer(http.Dir("assets/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    fmt.Println("server running...")
    http.ListenAndServe(":80", nil)
}