package main

import (
	"context"
	"encoding/json"
	"fmt"
	api_endpoints "github.com/tgwilliams/simple-search-ui/internal/api-endpoints"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
)

func main() {
	r := api_endpoints.CreateRouter()

	r.Run(":8080")
}

func loadData(fileName string, collection *mongo.Collection) {
	fmt.Println("Loading Sample Data!")
	var data []map[string]interface{}
	dataBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Couldn't load example data file")
		panic(err.Error())
	}

	json.Unmarshal(dataBytes, &data)
	for _, object := range data {
		fmt.Printf("Object %+v\n", object)
		collection.InsertOne(context.TODO(), object)
	}

	//collection.InsertMany(context.TODO(), string(dataBytes))
}