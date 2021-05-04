package document_connector

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Field struct {
	ID primitive.ObjectID `bson:"_id"`
	CombinedName string `bson:"combined_name"`
	FieldName string `bson:"field_name"`
	Schema string `bson:"schema"`
	Type string `bson:"type"`
}

func (dc *DocClientConfig) GetFields() []Field {
	collection := dc.Client.Database("SCHEMAS").Collection("schema")
	cursor, _ := collection.Find(dc.QueryCtx, bson.D{})

	var results []Field
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	return results
}

func (dc *DocClientConfig) GetFieldsWithParams(queryList [][]string) []Field {
	collection := dc.Client.Database("SCHEMAS").Collection("schema")

	fmt.Printf("Query List: %+v\n",queryList)
	var filter bson.D
	for _, queryObj := range(queryList) {
		regexPattern := fmt.Sprintf(".*%s.*", queryObj[1])
		filter = append(filter, bson.E{Key:queryObj[0], Value: bson.D{
			{"$regex", primitive.Regex{Pattern:regexPattern, Options:"i"}},
		}})
	}

	cursor, err := collection.Find(dc.QueryCtx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var results []Field
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	if len(results) > 0 {
		return results
	} else {
		return []Field{}
	}
}
