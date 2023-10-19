package main

import (
	"github.com/gocolly/colly"
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

// route to handle post requests for scraped data
func updateHouses(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
		return
	}
	limit, _ := strconv.Atoi(r.PostFormValue("limit"))
	url := r.PostFormValue("url")

	c, err := Scrape(url, "houses.csv", limit)
	var tpl = template.Must(template.ParseFiles("templates/results.html"))

	c.OnScraped(func(response *colly.Response) {
		err = tpl.Execute(w, "Data scraped successfully")
		if err != nil {
			panic(err)
			return
		}
	})
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", getAllHouses)
	router.HandleFunc("/scrape", updateHouses)

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
