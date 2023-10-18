package main

import (
	"github.com/gocolly/colly"
	"os"
	"testing"
)

func TestScrapeWebsite(t *testing.T) {
	c, err := Scrape("https://www.buyrentkenya.com/houses-for-rent/", "tests.csv", 2)

	if err != nil {
		t.Fail()
	}
	c.OnError(func(_ *colly.Response, err error) {
		t.Fail()
	})
	c.OnResponse(func(response *colly.Response) {
		if response.Request.URL.Path != "https://www.buyrentkenya.com/houses-for-rent/" {
			t.Fail()
		}
	})
	c.OnScraped(func(r *colly.Response) {
		file, err := os.Open("test.csv")
		if err != nil {
			t.Fail()
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				panic(err)
				return
			}
		}(file)
	})
}

//=== RUN   TestScrapeWebsite
//--- PASS: TestScrapeWebsite (3.67s)
//PASS
