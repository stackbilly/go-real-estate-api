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

func main() {
	c := colly.NewCollector()

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

	err := c.Visit("https://www.buyrentkenya.com/houses-for-rent")
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
