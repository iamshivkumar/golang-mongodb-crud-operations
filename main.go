package main

import (
	"context"
	"github.com/shivkumar123g/mongodb/src"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	// err := AddNumber(Number{
	// 	ID:    primitive.NewObjectID(),
	// 	Name:  "a",
	// 	Value: 1.0,
	// })

	// err := AddNumbers([]Number{
	// 	Number{
	// 		ID: primitive.NewObjectID(),
	// 		Name: "pi",
	// 		Value: 3.14,
	// 	},
	// 	Number{
	// 		ID: primitive.NewObjectID(),
	// 		Name: "a",
	// 		Value: 1.0,
	// 	},
	// })

	// number, err := GetNumberByName("pi")
	// log.Println(number)

	// numbers,err := GetAllNumbers()
	// log.Println(numbers)

	// err:= DeleteOneNumber("a")

	// err := DeleteAllNumbers()

	err := UpdateNumber("pi")
	if err != nil {
		log.Print(err)
	}
}

var db string = "testing"
var collection string = "numbers"

//Number is
type Number struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name"`
	Value float64            `bson:"value"`
}

//AddNumber is
func AddNumber(number Number) error {
	client, err := src.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(db).Collection(collection)
	_, err = collection.InsertOne(context.TODO(), number)
	if err != nil {
		return err
	}
	return nil
}

//AddNumbers is
func AddNumbers(list []Number) error {
	insertableList := make([]interface{}, len(list))
	for i, v := range list {
		insertableList[i] = v
	}
	client, err := src.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(db).Collection(collection)
	_, err = collection.InsertMany(context.TODO(), insertableList)
	if err != nil {
		return err
	}
	return nil
}

//GetNumberByName is
func GetNumberByName(name string) (Number, error) {
	result := Number{}
	filter := bson.D{primitive.E{Key: "name", Value: "pi"}}
	client, err := src.GetMongoClient()
	if err != nil {
		return result, err
	}
	collection := client.Database(db).Collection(collection)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return result, err
	}
	return result, nil
}

//GetAllNumbers is
func GetAllNumbers() ([]Number, error) {
	filter := bson.D{{}}
	var numbers []Number
	client, err := src.GetMongoClient()
	if err != nil {
		return numbers, err
	}
	collection := client.Database(db).Collection(collection)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return numbers, findError
	}
	for cur.Next(context.TODO()) {
		var t Number
		err := cur.Decode(&t)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, t)
	}
	cur.Close(context.TODO())
	if len(numbers) == 0 {
		return numbers, mongo.ErrNoDocuments
	}
	return numbers, nil
}

//DeleteOneNumber is
func DeleteOneNumber(name string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	client, err := src.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(db).Collection(collection)
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

//DeleteAllNumbers is
func DeleteAllNumbers() error {
	selector := bson.D{{}}
	client, err := src.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(db).Collection(collection)
	_, err = collection.DeleteMany(context.TODO(), selector)
	if err != nil {
		return err
	}
	return nil
}

// UpdateNumber is
func UpdateNumber(name string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "value", Value: "3.142"},
	}}}
	client, err := src.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(db).Collection(collection)
	_, err = collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return err
	}
	return nil
}
