package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	port := 8247
	http.HandleFunc("/", handler)
	// I KNOW
	log.Println("Server started on port " + strconv.Itoa(port))
	err := http.ListenAndServe("localhost:"+strconv.Itoa(port), nil)
	if err != nil {
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "There is an error occurred reading request: ", http.StatusInternalServerError)
		}
		log.Println("Request body: " + string(body))
	}
	randomInt := randIntInRange(1, 100)

	_, err := w.Write([]byte(strconv.Itoa(randomInt)))
	log.Println("Random number sent:", randomInt)
	if err != nil {
		return
	}
}

func randIntInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}
