package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	dir := "."
	if len(os.Args) >= 2 {
		dir = os.Args[1]
	}
	const addr = "localhost:8080"
	log.Printf("Run Web Server on http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir(dir))))
}
