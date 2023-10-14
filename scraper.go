package main

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"os"
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

func main() {
	c := colly.NewCollector()

	c.OnHTML(".listing-card", func(e *colly.HTMLElement) {
		house := House{}
		house.Name = e.ChildText(".hide-title")
		house.Location = e.ChildText("p")
		house.Description = e.ChildText(".text-md")
		house.Price = e.ChildText("p")
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
