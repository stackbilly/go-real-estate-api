package main

import (
	"testing"
)

func TestScrapeWebsite(t *testing.T) {
	if err != nil {
		panic(err)
		return
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1 test if function scrapes and writes to csv",
			args: args{filename: "test.csv"},
			want: *file,
		},
	}
}
