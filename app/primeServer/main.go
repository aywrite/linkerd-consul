package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// TODO
// Use a routing library
// Check HTTP method
// break handlers out to separate file
// api documentation
// landing page

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping/", pingHandler)
	mux.HandleFunc("/prime/", primeHandler)
	http.ListenAndServe(":8080", mux)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	switch status := os.Getenv("HEALTH"); status {
	case "critical":
		w.WriteHeader(http.StatusInternalServerError)
	case "warning":
		w.WriteHeader(http.StatusTooManyRequests)
	case "passing":
		w.WriteHeader(http.StatusOK)
	default:
		panic(fmt.Sprintf("Invalid health status %s", status))
	}
	io.WriteString(w, "pong")
}

func primeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	id := strings.TrimPrefix(r.URL.Path, "/prime/")
	id = strings.TrimSuffix(id, "/")
	num, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	next := nextPrime(num)
	resp := fmt.Sprintf("The next prime after %v is %v", num, next)
	io.WriteString(w, resp)
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1 // negative numbers aren't prime
}

func nextPrime(num int) int {
	for i := num + 1; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}
