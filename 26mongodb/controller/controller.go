package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mymongodb/model"
	"net/http"

	"github.com/gorilla/mux"
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

//delete 1 record

func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id":id}

	deleteCount,err :=collection.DeleteOne(context.Background(),filter)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count: ",deleteCount)

}

//delete all records from mongodb
func deleteAllMovies() int64{
	 deleteResult, err := collection.DeleteMany(context.Background(),bson.D{{}},nil)
	 if err != nil{
		log.Fatal(err)
	}
	fmt.Println("No of movies deleted",deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}
//get all movies from mongodb

func getAllMovies()[]primitive.M{  
	cur,err  := collection.Find(context.Background(),bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M 
		err := cur.Decode(&movie)
		if err != nil{
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

//Acutal controller - file

func GetMyAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie (w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched (w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}