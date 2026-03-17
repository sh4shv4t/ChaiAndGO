package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/sh4shv4t/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func ConnectDB() {
	connectionString := os.Getenv("MONGODB_URI")
	if connectionString == "" {
		log.Fatal("MONGODB_URI is not set. Add it to .env or your shell environment")
	}

	dbName := getEnv("MONGODB_DATABASE", "netflix")
	colName := getEnv("MONGODB_COLLECTION", "watchlist")

	// Client options
	clientOptions := options.Client().ApplyURI(connectionString)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection successful")

	collection = client.Database(dbName).Collection(colName)

	// Collection instance
	fmt.Println("Collection instance ready")
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// Mongo helpers
func insertOneMovie(movie model.Netflix) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	inserted, err := collection.InsertOne(ctx, movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie with ID:", inserted.InsertedID)
}

func updateOneMovie(movieID primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": movieID}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count:", result.ModifiedCount)
}

func deleteOneMovie(movieID primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": movieID}
	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count:", deleteResult.DeletedCount)
}

func deleteAllMovies() int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	deleteResult, err := collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted movies count:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func getAllMovies() []primitive.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var movies []primitive.M
	for cursor.Next(ctx) {
		var movie bson.M
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return movies
}

func getOneMovie(movieID primitive.ObjectID) (primitive.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var movie bson.M
	err := collection.FindOne(ctx, bson.M{"_id": movieID}).Decode(&movie)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

// HTTP handlers
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getAllMovies())
}

func GetOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	movieID := params["id"]
	if movieID == "" {
		http.Error(w, "missing movie id", http.StatusBadRequest)
		return
	}

	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		http.Error(w, "invalid movie id", http.StatusBadRequest)
		return
	}

	movie, err := getOneMovie(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, "movie not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to fetch movie", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movie)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie model.Netflix
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	movieID := params["id"]
	if movieID == "" {
		http.Error(w, "missing movie id", http.StatusBadRequest)
		return
	}
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		http.Error(w, "invalid movie id", http.StatusBadRequest)
		return
	}

	updateOneMovie(id)
	json.NewEncoder(w).Encode(map[string]string{"message": "movie marked as watched"})
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	movieID := params["id"]
	if movieID == "" {
		http.Error(w, "missing movie id", http.StatusBadRequest)
		return
	}
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		http.Error(w, "invalid movie id", http.StatusBadRequest)
		return
	}

	deleteOneMovie(id)
	json.NewEncoder(w).Encode(map[string]string{"message": "movie deleted successfully"})
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(map[string]interface{}{"deletedCount": count})
}
