package main

import (
	mongoimport "github.com/Livingstone-Billy/mongo-import"
	"testing"
)

func TestScrapeWebsite(t *testing.T) {
	records, err := mongoimport.CSVReader("test.csv")
	if err != nil {
		panic(err)
	}
	type args struct {
		filename string
		url      string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1 test if function scrapes and writes to csv",
			args: args{filename: "test.csv", url: "https://www.buyrentkenya.com/houses-for-rent/"},
			want: len(records),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ScrapeWebsite(tt.args.url, tt.args.filename)
			if err != nil {
				t.Fail()
				panic(err)
			}
			assert.
		})
	}
}
