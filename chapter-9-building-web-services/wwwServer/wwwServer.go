package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>", t)

	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments)!=1 {
		PORT = ":" + arguments[1]
	}

	fmt.Println("Using port number: ", PORT)

	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe(PORT, nil)
	if err!=nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}