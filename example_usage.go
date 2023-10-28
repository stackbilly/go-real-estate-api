package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"net/http"
)

//define struct to hold json data

type Houses struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"Name" json:"Name"`
	Description string        `bson:"Description" json:"Description"`
	Location    string        `bson:"Location" json:"Location"`
	Price       string        `bson:"Price" json:"Price"`
	Url         string        `bson:"Url" json:"Url"`
	Image       string        `bson:"Image" json:"Image"`
}

// HandleHouseData sample http request to make API call
func HandleHouseData(w http.ResponseWriter, _ *http.Request) {
	resp, err := http.Get("http://localhost:8080/api/houses")
	if err != nil {
		panic(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
			return
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Unexpected status code", resp.StatusCode)
	}
	var house Houses
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		return
	}
	houseJson := data
	if err = json.Unmarshal(houseJson, &house); err != nil {
		panic(err)
		return
	}
	//do what you want with the data
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(houseJson)
	if err != nil {
		return
	}
}
