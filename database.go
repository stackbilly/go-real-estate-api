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

// ImportToDB ,write csv file to db
func ImportToDB() (int, error) {
	session, err := getSession()
	if err != nil {
		panic(err)
		return 0, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

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

// Retrieve  all houses from database
func Retrieve() ([]byte, error) {
	session, err := getSession()
	if err != nil {
		panic(err)
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
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

// InsertDB inserts scraped data to DB without reading from csv
func InsertDB(houses []House) (int, error) {
	session, err := getSession()
	if err != nil {
		panic(err)
		return 0, err
	}
	defer session.Close()
	col := session.DB("estate").C("houses")
	c, _ := col.Count()
	if c > 0 {
		err = col.DropCollection()
		if err != nil {
			panic(err)
			return 0, err
		}
	}
	for _, data := range houses {
		err = col.Insert(data)
		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		panic(err)
		return 0, err
	}
	c, err = col.Count()
	if err != nil {
		panic(err)
		return 0, err
	}
	return c, nil
}
