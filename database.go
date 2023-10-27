package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func getClient() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	password := os.Getenv("PASSWORD")
	uri := fmt.Sprintf("mongodb+srv://LaplaceBilly:%s@cluster0.vhmgz.mongodb.net/?retryWrites=true&w=majority", password)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// ImportToDB ,write csv file to db
//func ImportToDB() (int, error) {
//	session, err := getSession()
//	if err != nil {
//		panic(err)
//		return 0, err
//	}
//	defer session.Close()
//	session.SetMode(mgo.Monotonic, true)
//
//	collection := session.DB("estate").C("houses")
//	records, err := mongoimport.CSVReader("houses.csv")
//	if err != nil {
//		panic(err)
//		return 0, err
//	}
//	start := time.Now()
//	count := mongoimport.CSVImport(collection, records, 1, len(records))
//	fmt.Printf("Inserted %d records in %s seconds", count, time.Since(start))
//
//	return count, nil
//}

// Retrieve  all houses from database
func Retrieve() ([]byte, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
			return
		}
	}()
	collection := client.Database("estate").Collection("houses")
	ctx := context.TODO()

	filter := bson.M{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			panic(err)
			return
		}
	}(cursor, ctx)
	jsonData, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func RetrieveLimit(limit int) ([]House, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
			return
		}
	}()
	ctx := context.TODO()

	collection := client.Database("estate").Collection("houses")

	filter := bson.M{}
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			panic(err)
			return
		}
	}(cursor, ctx)

	var results []House
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// SaveToDatabase inserts scraped data to DB without reading from csv
func SaveToDatabase(houses []House) (int, error) {
	client, err := getClient()
	if err != nil {
		return 0, err
	}
	defer func() {
		err := client.Disconnect(context.TODO())
		if err != nil {
			panic(err)
			return
		}
	}()

	ctx := context.TODO()

	collection := client.Database("estate").Collection("houses")
	filter := bson.M{}
	_, err = collection.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	documents := make([]interface{}, len(houses))
	for i, data := range houses {
		documents[i] = data
	}
	_, err = collection.InsertMany(ctx, documents)
	if err != nil {
		return 0, err
	}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
