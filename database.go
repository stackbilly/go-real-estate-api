package main

import (
	"encoding/json"
	"fmt"
	mongoimport "github.com/Livingstone-Billy/mongo-import"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func getSession() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
		return nil, err
	}
	return session.Copy(), err
}

func ImportHouseToDatabase() (int, error) {
	session, err := getSession()
	if err != nil {
		panic(err)
		return 0, err
	}
	defer session.Close()

	collection := session.DB("estate").C("houses")
	records, err := mongoimport.CSVReader("houses.csv")
	if err != nil {
		panic(err)
		return 0, err
	}
	start := time.Now()
	count := mongoimport.CSVImport(collection, records, 1, len(records))
	fmt.Printf("Inserted %d records in %s seconds", count, time.Since(start))

	return count, nil
}

func RetrieveAllHouses() ([]byte, error) {
	session, err := getSession()
	if err != nil {
		panic(err)
		return nil, err
	}
	defer session.Close()
	query := bson.M{}

	var results []bson.M
	collection := session.DB("estate").C("houses")

	err = collection.Find(query).All(&results)
	if err != nil {
		panic(err)
		return nil, err
	}
	jsonData, err := json.Marshal(results)
	return jsonData, nil
}
