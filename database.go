package main

import (
	"encoding/json"
	"fmt"
	mongoimport "github.com/Livingstone-Billy/mongo-import"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"
)

func getSession() (*mgo.Session, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	url := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.vhmgz.mongodb.net/estate", username, password)
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Printf("Connection err:\n%s", err)
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

func RetrieveLimit(limit int) ([]House, error) {
	session, err := getSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	coll := session.DB("estate").C("houses")

	query := coll.Find(nil).Limit(limit)

	iter := query.Iter()
	var houses []House

	var result House
	for iter.Next(&result) {

		houses = append(houses, result)
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return houses, nil
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
