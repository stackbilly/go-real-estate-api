package main

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestClient(t *testing.T) {
	client, err := getClient()
	if err != nil {
		fmt.Printf("\nFail to get client %s", err)
		t.Fail()
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			fmt.Printf("\nFail to disconnect client %s", err)
			return
		}
	}()
	//sending a ping to check successful connection
	var result bson.M
	if err = client.Database("admin").RunCommand(context.TODO(),
		bson.D{{"ping", 1}}).Decode(&result); err != nil {
		fmt.Printf("\nFail to ping %s", err)
		t.Fail()
	} else {
		t.Log("\nSuccessfully connected to mongodb atlas")
	}
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

func TestSaveToDatabase(t *testing.T) {
	var houses []House
	houses = append(houses, House{
		Name:        "4 Bed House in Nyari",
		Description: "4 Bed House in Nyari",
		Location:    "Nyari, Westlands",
		Price:       "Price not communicated",
		Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
		Image:       "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
	},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Image:       "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
		},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Image:       "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
		},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Image:       "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
		},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Image:       "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
		},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Image:       "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
		},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Image:       "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
		},
		House{
			Name:        "4 Bed House in Nyari",
			Description: "4 Bed House in Nyari",
			Location:    "Nyari, Westlands",
			Price:       "Price not communicated",
			Url:         "/listings/4-bedroom-house-for-rent-nyari-3623660",
			Image:       "https://i.roamcdn.net/prop/brk/listing-thumb-376w/c18df8c33939424def6ea66c405232c6/-/prod-property-core-backend-media-brk/5520842/6deae430-b3e5-4ab1-90e7-966f30cc791a.jpg",
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
			got, err := SaveToDatabase(houses)
			if err != nil {
				t.Fail()
			}
			assert.Equal(t, tt.want, got, "Unexpected results got on test run")
		})
	}
}

func TestRetrieveLimit(t *testing.T) {
	docs, err := RetrieveLimit(10)
	if err != nil {
		t.Fail()
	}
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test if len(docs) retrieved = limit",
			args: args{limit: 10},
			want: len(docs),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := RetrieveLimit(tt.args.limit)
			got := len(data)
			if err != nil {
				t.Fail()
			}
			assert.Equal(t, tt.want, got, "Got unexpected value")
		})
	}
}
