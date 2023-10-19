package main

import (
	mongoimport "github.com/Livingstone-Billy/mongo-import"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImportToDB(t *testing.T) {
	records, err := mongoimport.CSVReader("houses.csv")
	if err != nil {
		panic(err)
		return
	}
	want := len(records) - 1

	got, err := ImportToDB()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, want, got, "Expected number of records in the database")
}

func TestRetrieve(t *testing.T) {
	jsonData, err := Retrieve()
	if err != nil {
		t.Fail()
	}
	if jsonData == nil {
		t.Fail()
	}
}

// test insert
func TestInsertDB(t *testing.T) {
	var houses []House
	houses = append(houses, House{
		Name:        "4 Bed House in Nyari",
		Description: "4 Bed House in Nyari",
		Location:    "Nyari, Westlands",
		Price:       "Price not communicated",
		Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
		Img:         "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
	},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Img:         "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
		},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Img:         "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
		},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Img:         "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
		},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Img:         "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
		},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Img:         "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
		},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Img:         "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
		},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Img:         "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
		},
	)
	type args struct {
		house []House
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Match len of slice equal len of inserted records",
			args: args{house: houses},
			want: len(houses),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InsertDB(houses)
			if err != nil {
				t.Fail()
			}
			assert.Equal(t, tt.want, got, "Unexpected results got on test run")
		})
	}
}
