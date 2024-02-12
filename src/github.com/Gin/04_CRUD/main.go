package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type manager struct {
	client 		*mongo.Client
	ctx 		context.Context
	cancel 		context.CancelFunc
}

var Mgr manager

type Manager interface {

	Insert(interface{}) error
	GetAll() ([]User, error)
	DeleteData(primitive.ObjectID) error
	UpdateData(User) error

}

func DBConnection() {

	connectionURI := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(connectionURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	
	client, err := mongo.Connect(ctx, clientOptions)
	ErrorFunc(err)

	err = client.Ping(ctx, nil)
	ErrorFunc(err)

	fmt.Println("Connected...!!")
	Mgr = manager{client: client, ctx: ctx, cancel: cancel} 

}

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()

	defer func() {
		
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func ErrorFunc(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	DBConnection()
}

type User struct {
	ID		primitive.ObjectID	`bson:"_id,omitempty"`
	Name	string				`bson:"name"`
	Email	string				`bson:"email"`
}

func (m *manager) Insert(data interface{}) error {
	
	collection := m.client.Database("demoDB").Collection("demoCollection")
	result, err := collection.InsertOne(m.ctx, data)
	// ErrorFunc(err)

	fmt.Println(result.InsertedID)
	return err
}

func (m *manager)GetAll() (data []User, err error) {

	collection := m.client.Database("demoDB").Collection("demoCollection")

	// Pass this options to the Find Method
	findOptions := options.Find()

	cur, _ := collection.Find(m.ctx, bson.M{}, findOptions)
	for cur.Next(m.ctx) {
		var d User
		err := cur.Decode(&d)
		ErrorFunc(err)
		data = append(data, d)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(m.ctx)

	return data, nil

}

func (m *manager) DeleteData(id primitive.ObjectID) error {

	collection := m.client.Database("demoDB").Collection("demoCollection")

	filter := bson.D{{"_id", id}}
	_, err := collection.DeleteOne(m.ctx, filter)
	return err

}

func (m *manager) UpdateData(data User) error {

	collection := m.client.Database("demoDB").Collection("demoCollection")

	filter := bson.D{{"_id", data.ID}}
	update := bson.D{{"$set", data}}

	_, err := collection.UpdateOne(m.ctx, filter, update)

	return err

}

func main() {

	// Insert Record to mongodb
	u := User{Name: "abc", Email: "abc@abc.abc"}
	err := Mgr.Insert(u)
	ErrorFunc(err)

	// get all records from db
	data, err := Mgr.GetAll()
	ErrorFunc(err)
	fmt.Println(data)

	// delete record from db
	id := ""
	objectId, err := primitive.ObjectIDFromHex(id)
	ErrorFunc(err)
	err = Mgr.DeleteData(objectId)

	// Update
	objectId, err = primitive.ObjectIDFromHex(id)
	ErrorFunc(err)

	u.ID = objectId
	u.Name = "test"
	u.Email = "abc@abc.abc"
	err = Mgr.UpdateData(u)
	ErrorFunc(err)

}