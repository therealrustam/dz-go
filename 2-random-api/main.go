package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

func getRandom(w http.ResponseWriter, r *http.Request) {
	randomNumber := rand.Intn(6) + 1
	w.Write([]byte(strconv.Itoa(randomNumber)))

}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/random", getRandom)
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}

}
