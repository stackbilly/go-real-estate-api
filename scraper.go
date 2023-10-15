package main

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"os"
	"strings"
)

type House struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Price       string `json:"price"`
	Url         string `json:"url"`
	Img         string `json:"img"`
}

var houses []House

func splitString(substrings string, index int) string {
	subs := strings.Split(substrings, "\n")
	return subs[index]
}

func contains(slice []string, val string) bool {
	for _, key := range slice {
		if val == key {
			return true
		} else {
			return false
		}
	}
	return false
}

func main() {
	c := colly.NewCollector()

	var pagesToScrape []string
	pageToScrape := "https://www.buyrentkenya.com/houses-for-rent?page=1"
	pagesDiscovered := []string{pageToScrape}
	//current iteration
	i := 1
	//max pages to scrap
	limit := 5
	c.OnHTML("li.page-item", func(e *colly.HTMLElement) {
		newPaginationLink := e.Attr("href")

		//if page discovered is new
		if !contains(pagesToScrape, newPaginationLink) {
			//if the page discovered should be scraped
			if !contains(pagesDiscovered, newPaginationLink) {
				pagesToScrape = append(pagesToScrape, newPaginationLink)
			}
			pagesDiscovered = append(pagesDiscovered, newPaginationLink)
		}
	})
	c.OnHTML(".listing-card", func(e *colly.HTMLElement) {
		house := House{}
		house.Name = splitString(e.ChildText("span"), 1)
		house.Location = e.ChildText(".ml-1")
		house.Description = splitString(e.ChildText(".mb-3"), 0)
		house.Price = splitString(e.ChildText(".text-xl"), 0)
		house.Url = e.ChildAttr("a", "href")
		house.Img = e.ChildAttr("img", "src")

		houses = append(houses, house)
	})

	c.OnScraped(func(response *colly.Response) {
		if len(pagesToScrape) != 0 && i < limit {
			//get current page to scrape and remove it from list
			pageToScrape = pagesToScrape[0]
			pagesToScrape = pagesToScrape[1:]
			i++
			err := c.Visit(pageToScrape)
			if err != nil {
				panic(err)
				return
			}
		}
	})

	err := c.Visit(pageToScrape)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("houses.csv")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
			return
		}
	}()
	writer := csv.NewWriter(file)

	headers := []string{
		"Name",
		"Description",
		"Location",
		"Price",
		"Url",
		"Image",
	}
	err = writer.Write(headers)
	if err != nil {
		panic(err)
	}
	for _, house := range houses {
		record := []string{
			house.Name,
			house.Description,
			house.Location,
			house.Price,
			house.Url,
			house.Img,
		}
		err = writer.Write(record)
		if err != nil {
			panic(err)
		}
	}
	defer writer.Flush()
}
