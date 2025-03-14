package controller

import (
	"context"
	"fmt"
	"log"
	"mymongodb/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString ="mongodb+srv://pratim:zM2tgQMH5KVVgolA@golangcluster.uoaen.mongodb.net/?retryWrites=true&w=majority&appName=golangCluster"

const dbName ="netflix"
const collectionName ="watchlist"

var collection *mongo.Collection

///connect with mongodb

func init(){
	clientOptions := options.Client().ApplyURI(connectionString)
	
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDb connection success")
	collection = client.Database(dbName).Collection(collectionName)

	///collection instance

	fmt.Println("Collection instance is ready")
}
//Mongo Db helpers


//insert 1 record

func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(),movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)

}
//update 1 record

func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}

	result , err := collection.UpdateOne(context.Background(),filter,update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modied count: ", result.ModifiedCount)
}