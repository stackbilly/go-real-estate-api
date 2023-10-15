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

func WriteToCSV(filename string, houses []House) error {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
		return err
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
		return err
	}

	for _, house := range houses {
		records := []string{
			house.Name,
			house.Description,
			house.Location,
			house.Price,
			house.Url,
			house.Img,
		}
		err = writer.Write(records)
		if err != nil {
			panic(err)
			return err
		}
	}
	defer writer.Flush()
	return nil
}

func main() {
	c := colly.NewCollector(
		colly.Async(true),
	)

	// setting  User-Agent header
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.DisableCookies()
	c.AllowURLRevisit = true
	err := c.Limit(&colly.LimitRule{
		Parallelism: 4,
		DomainGlob:  "*",
	})
	if err != nil {
		panic(err)
		return
	}

	var pagesToScrape []string
	pageToScrape := "https://www.buyrentkenya.com/houses-for-rent/"
	pagesDiscovered := []string{pageToScrape}
	//current iteration
	i := 1
	//max pages to scrap
	limit := 110
	c.OnHTML("a.relative", func(e *colly.HTMLElement) {
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
		house.Name = formatName(e.ChildText("span"))
		house.Location = e.ChildText(".ml-1")
		house.Description = splitString(e.ChildText(".mb-3"))
		house.Price = splitString(e.ChildText(".text-xl"))
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
	err = c.Visit(pageToScrape)
	if err != nil {
		panic(err)
	}
	c.Wait()
	err = WriteToCSV("houses.csv", houses)
	if err != nil {
		panic(err)
		return
	}
}
