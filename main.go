package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var port string

func getenv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Env of %s is empty\n", key)
	}
	return val
}

func init() {
	port = getenv("PORT")
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is homepage.\n")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	log.Printf("Listen on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
