package main

import (
	mongoimport "github.com/Livingstone-Billy/mongo-import"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImportHouseToDatabase(t *testing.T) {
	records, err := mongoimport.CSVReader("houses.csv")
	if err != nil {
		panic(err)
		return
	}
	want := len(records) - 1

	got, err := ImportHouseToDatabase()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, want, got, "Expected number of records in the database")
}

//=== RUN   TestImportHouseToDatabase
//Inserted 1578 records in 2.8993034s seconds--- PASS: TestImportHouseToDatabase (3.23s)
//PASS

func TestRetrieveAllHouses(t *testing.T) {
	jsonData, err := RetrieveAllHouses()
	if err != nil {
		t.Fail()
	}
	if jsonData == nil {
		t.Fail()
	}
}

/*
=== RUN   TestRetrieveAllHouses
--- PASS: TestRetrieveAllHouses (0.05s)
PASS
*/
