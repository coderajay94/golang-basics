package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/coderajay94/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName             string = "netflix"
	collectionName     string = "watchlist"
	dbConnectionString string = "mongodb+srv://ajakumar:admin@cluster0.hecejb8.mongodb.net/"
)

var collection *mongo.Collection

//connect with mongo db

// only runs one time, only when application is starting
func init() {

	//client option
	clientOptions := options.Client().ApplyURI(dbConnectionString)

	//connect to mongo db
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("error occured while connecting to mongodb database")
		log.Fatal(err)
	}

	fmt.Println("client connected successfully")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("collection instance is ready")
}

// mongo db helpers
// insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("insterd 1 movie with id:", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) {

	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		fmt.Println("error occured while converting mongo id in updateOneMovie")
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, errr := collection.UpdateOne(context.Background(), filter, update)
	if errr != nil {
		log.Fatal(errr)
	}

	fmt.Println("modiffied Count:", result.ModifiedCount)
}

//delete 1 record

func deleteOneMovie(movieId string) {

	//convert the string id to mongo db understandable id
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		fmt.Println("error occured while converting mongo id in deleteOneMovie")
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	deletedResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println("error occured while deleting one record in deleteOneMovie")
		log.Fatal(err)
	}
	fmt.Println("record deleted with count ", deletedResult.DeletedCount)
}

// delete all movies
func deleteAllMovies() int {

	//we can user bson.D and bson.M - its just D provides ordered list,
	//also {{}} empty filter means all the records will be selected
	filter := bson.D{{}}
	deletedResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		fmt.Println("error occured while deleting all movies")
		log.Fatal(err)
	}

	fmt.Println("all movies deleted with count ", deletedResult.DeletedCount)
	return int(deletedResult.DeletedCount)
}

//get all movies from database

func getAllMovies() []primitive.M {

	//when you fetch elements from mongo one cursor will be returned
	//use that cursor to get all elements
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		fmt.Println("error while fetching all movies from database: ")
		log.Fatal(err)
	}

	//do cursor.next and map that to an object, if valid add into your list or returing objects

	var movies []primitive.M

	for cur.Next(context.Background()) {

		var movie bson.M
		err := cur.Decode(&movie)

		if err != nil {
			fmt.Println("error occured while decoding the cursor next object")
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

// create handler functions
func GetAllMyMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	movies := getAllMovies()
	json.NewEncoder(w).Encode(movies)
}

// create new movie
func CreateNewMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
	fmt.Println("Movie created successfully")
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
	fmt.Println("movie is update to watch successfully")
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
	fmt.Println("Deleted one Movie: ", params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()

	json.NewEncoder(w).Encode(count)
	fmt.Println("All Movies are deleted successfully...")
}
