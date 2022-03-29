package quickstart

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type Person struct {
	Name string
	Age  int
}

func Sample1() {

	uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

	client := getClient(uri)

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	//ping(client)

	//insert(client)

	findAll(client)

}

func getClient(uri string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}

func ping(client *mongo.Client) {

	err := client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}

}

func findWithCriteria(client *mongo.Client) {

	coll := client.Database("belajardb").Collection("person")
	var result Person
	err := coll.FindOne(context.TODO(), bson.M{"age": 12}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found \n")
		return
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func insert(client *mongo.Client) {

	person := Person{
		Name: "Omar",
		Age:  31,
	}

	coll := client.Database("belajardb").Collection("person")
	result, err := coll.InsertOne(context.TODO(), person)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", result.InsertedID)

}

func findAll(client *mongo.Client) {

	coll := client.Database("belajardb").Collection("person")

	//cur, err := coll.Find(context.TODO(), bson.M{"age": bson.M{"$gt": 17}})

	//cur, err := coll.Find(context.TODO(), bson.M{
	//	"age": bson.M{
	//		"$in": bson.A{33, 31},
	//	},
	//})

	cur, err := coll.Find(context.TODO(), bson.D{{
		Key: "age",
		Value: bson.D{
			{
				Key:   "$in",
				Value: bson.A{33, 31},
			},
		},
	}})

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found \n")
		return
	}
	if err != nil {
		panic(err)
	}

	for cur.Next(context.TODO()) {
		var result Person
		if err := cur.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", result)
	}

}
