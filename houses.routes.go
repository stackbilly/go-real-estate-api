package main

import (
	"log"
	"net/http"
)

func getAllHouses(w http.ResponseWriter, r *http.Request) {
	house, err := RetrieveAllHouses()
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(house)
	if err != nil {
		panic(err)
		return
	}
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", getAllHouses)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
		return
	}
}
