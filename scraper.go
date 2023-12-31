package main

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

type House struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"Name" json:"Name"`
	Description string        `bson:"Description" json:"Description"`
	Location    string        `bson:"Location" json:"Location"`
	Price       string        `bson:"Price" json:"Price"`
	Url         string        `bson:"Url" json:"Url"`
	Image       string        `bson:"Image" json:"Image"`
}

var houses []House

func formatName(name string) string {
	var vSubs []string
	if strings.Compare(strings.Split(name, " ")[0], "Added") == 0 {
		subs := strings.Split(name, "\n")
		return subs[2]
	} else {
		vSubs = strings.Split(name, "\n")
		return vSubs[1]
	}
}

func splitString(substrings string) string {
	var subs []string

	subs = strings.Split(substrings, "\n")
	return subs[0]
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

func WriteToCSV(w http.ResponseWriter, houses []House) error {
	writer := csv.NewWriter(w)
	defer writer.Flush()

	headers := []string{
		"Name",
		"Description",
		"Location",
		"Price",
		"Url",
		"Image",
	}
	err := writer.Write(headers)
	if err != nil {
		panic(err)
		return err
	}

	for _, house := range houses {
		records := []string{
			house.Name,
			house.Description,
			house.Location,
			house.Price,
			house.Url,
			house.Image,
		}
		err = writer.Write(records)
		if err != nil {
			panic(err)
			return err
		}
	}

	return nil
}

func Scrape(url string, limit int) (colly.Collector, error) {
	//scraper configurations
	c := colly.NewCollector(colly.Async(true))

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.DisableCookies()
	c.AllowURLRevisit = true
	err := c.Limit(&colly.LimitRule{
		Parallelism: 4,
		DomainGlob:  "*",
	})
	if err != nil {
		panic(err)
	}
	var pagesToScrape []string
	pageToScrape := url
	pagesDiscovered := []string{pageToScrape}

	i := 1 //current iteration

	c.OnHTML("a.relative", func(e *colly.HTMLElement) {
		newPaginationLink := e.Attr("href")
		//if page discovered is new
		if !contains(pagesToScrape, newPaginationLink) {
			if !contains(pagesDiscovered, newPaginationLink) {
				pagesToScrape = append(pagesToScrape, newPaginationLink)
			}
			pagesDiscovered = append(pagesDiscovered, newPaginationLink)
		}
	})

	c.OnHTML(".listing-card", func(e *colly.HTMLElement) {
		house := House{}
		house.Name = formatName(e.ChildText("span"))
		house.Location = e.ChildText(".ml-1")
		house.Description = splitString(e.ChildText(".mb-3"))
		house.Price = splitString(e.ChildText(".text-xl"))
		house.Url = e.ChildAttr("a", "href")
		house.Image = e.ChildAttr("img", "src")

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

	err = c.Visit(pageToScrape)
	if err != nil {
		panic(err)
		return *c, err
	}
	c.Wait()

	//err = WriteToCSV(filename, houses)
	//if err != nil {
	//	panic(err)
	//	return *c, err
	//}

	_, err = SaveToDatabase(houses)
	if err != nil {
		panic(err)
		return *c, err
	}

	if err != nil {
		panic(err)
		return *c, err
	}
	return *c, nil
}
