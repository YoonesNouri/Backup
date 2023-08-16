package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// * Product describes an electronic product e.g. phone
type Product struct {
	ID          primitive.ObjectID `bson:"_id" bson:"_id"`
	Name        string             `json:"product_name" bson:"product_name"`
	Price       int                `json:"price" bson:"price"`
	Currency    string             `json:"currency" bson:"currency"`
	Quantity    string             `json:"quantity" bson:"quantity"`
	Discount    int                `json:"discount,omitempty" bson:"discount,omitempty"`
	Vendor      string             `json:"vendor" bson:"vendor"`
	Accessories []string           `json:"accessories,omitempty" bson:"accessories,omitempty"`
	SkuID       string             `json:"sku_id" bson:"sku_id"`
}

var iphone10 = Product{
	ID:          primitive.NewObjectID(),
	Name:        "iPhone 10",
	Price:       900,
	Currency:    "USD",
	Quantity:    "40",
	Vendor:      "Apple",
	Accessories: []string{"charger", "headset", "slot opener"},
	SkuID:       "1234",
}

var trimmer = Product{
	ID:          primitive.NewObjectID(),
	Name:        "easy trimmer",
	Price:       120,
	Currency:    "USD",
	Quantity:    "300",
	Vendor:      "Philips",
	Discount:    7,
	Accessories: []string{"charger", "comb", "blade set", "cleaning oil"},
	SkuID:       "2345",
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	//* primitive.ObjectID: primitive.NewObjectID() generates a unique MongoDB ObjectID.
	//* bson.A: bson.A{" value1", "value2" , …} creates an array with the provided values.
	//* bson.B: Not typically used directly; used internally for binary data.
	//* bson.D: bson.D{{"key", "value"}} creates an ordered document (slice of key-value pairs).
	//* bson.M: bson.M{"key": "value"} creates a map-style dynamic document.

	db := client.Database("tronics")
	collection := db.Collection("products")
	// res, err := collection.InsertOne(context.Background(), trimmer)
	// res, err := collection.InsertOne(context.Background(), bson.D{
	// 	{"name", "eric"},
	// 	{"surname", "cartman"},
	// 	{"hobbies", bson.A{"videogame", "alexa", "kfc"}},
	// })
	res, err := collection.InsertOne(context.Background(), bson.M{
		"name":    "eric",
		"surname": "cartman",
		"hobbies": bson.A{"videogame", "alexa", "kfc"},
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
}
