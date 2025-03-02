package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongo/model"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionn *mongo.Collection

const collectionName = "watchlist"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables.")
	}

	connectionString := os.Getenv("MONGO_URI")
	fmt.Println(connectionString)
	dbName := os.Getenv("DB_NAME")
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDB Connection Error: %v", err)
	}

	// Ping the database to confirm a successful connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("MongoDB Ping Failed: %v", err)
	} else {
		fmt.Println("MongoDB Connection Success")
	}

	collectionn = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection reference is ready")
}

// mongodb helpers  - file

// insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collectionn.InsertOne(context.Background(), movie)
	ftlerr(err)
	fmt.Println("Inserted 1 movie in db wiith id: ", inserted.InsertedID)

}

func updateOneRecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collectionn.UpdateOne(context.Background(), filter, update)
	ftlerr(err)
	fmt.Println("Updated 1 movie in db with id: ", result.UpsertedID)
}

func deleteOneRecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collectionn.DeleteOne(context.Background(), filter)
	ftlerr(err)
	fmt.Println("Deleted 1 movie in db with id: ", deleteCount)
}

// delete all records from mongodb
func deleteAllRecords() int64 {
	deleteCount, err := collectionn.DeleteMany(context.Background(), bson.D{}, nil)
	ftlerr(err)
	fmt.Println("Deleted all movies in db with count: ", deleteCount.DeletedCount)
	return deleteCount.DeletedCount
}

func getAllMovies() []model.Netflix {
	cur, err := collectionn.Find(context.Background(), bson.D{{}})
	ftlerr(err)
	defer cur.Close(context.Background())

	var movies []model.Netflix
	for cur.Next(context.Background()) {
		var movie model.Netflix
		err := cur.Decode(&movie) // Decode into `model.Netflix` struct
		ftlerr(err)
		movies = append(movies, movie)
	}
	return movies
}

// / actual controllers - file
func AllMovies(w http.ResponseWriter, r *http.Request) {
	movies := getAllMovies()
	w.Header().Set("Contennt-Type", "application/x-www-form-urlencode")
	json.NewEncoder(w).Encode(movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contennt-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

// function to mark movie as watched
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contennt-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contennt-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contennt-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteAllRecords()
	json.NewEncoder(w).Encode("All Movies deleted")
}

func ftlerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
