package main

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func getAllHouses(w http.ResponseWriter, _ *http.Request) {
	house, err := Retrieve()
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

func getScrape(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("templates/scrape.html"))
	var err error
	if r.Method == "GET" {
		err = tpl.Execute(w, nil)
	}
	if err != nil {
		http.Error(w, "Error navigating to scrape.html", http.StatusBadRequest)
	}
}

// route to handle post requests for scraped data
func updateHouses(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			panic(err)
			return
		}
		limit, _ := strconv.Atoi(r.PostFormValue("limit"))
		url := r.PostFormValue("url")

		c, err := Scrape(url, limit)

		c.OnScraped(func(response *colly.Response) {
			if err != nil {
				panic(err)
				return
			}
		})
	} else {
		http.Error(w, "Bad Method", http.StatusBadRequest)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := RetrieveLimit(5)
	if err != nil {
		panic(err)
		return
	}
	var tpl = template.Must(template.ParseFiles("templates/index.html"))
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Fail to execute template", http.StatusBadRequest)
	}
}

func getSingleHouse(w http.ResponseWriter, r *http.Request) {
	data, err := RetrieveLimit(2)
	if err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}
	j, err := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(j)
	if err != nil {
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/api/scrape", updateHouses).Methods("POST")
	router.HandleFunc("/api/scrape", getScrape).Methods("GET")
	router.HandleFunc("/api/houses", getAllHouses)
	router.HandleFunc("/api/house", getSingleHouse)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening....")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
