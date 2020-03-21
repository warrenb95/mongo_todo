package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/warrenb95/mongo_todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

// Creates a new Todo object in the database
func CreateTodoEndPoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	var todo models.Todo
	json.NewDecoder(req.Body).Decode(&todo)
	todo.TimeCreated = time.Now()

	collection := Client.Database("gotodo").Collection("todos")

	result, _ := collection.InsertOne(context.TODO(), todo)

	json.NewEncoder(res).Encode(result)
}

func GetAllTodosEndPoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	var todos []models.Todo
	collection := Client.Database("gotodo").Collection("todos")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var todo models.Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}
	if err := cursor.Err(); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(todos)
}

func GetTodoEndpoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	params := mux.Vars(req)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := Client.Database("gotodo").Collection("todos")

	var todo models.Todo
	err := collection.FindOne(context.TODO(), models.Todo{ID: id}).Decode(&todo)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(todo)
}

func DeleteTodoEndPoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := Client.Database("gotodo").Collection("todos")

	var todo models.Todo
	err := collection.FindOneAndDelete(context.TODO(), models.Todo{ID: id}).Decode(&todo)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(todo)
}

func UpdateTodoEndPoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := Client.Database("gotodo").Collection("todos")

	var todo models.Todo
	json.NewDecoder(req.Body).Decode(&todo)

	result, err := collection.UpdateOne(context.TODO(), models.Todo{ID: id}, bson.M{"$set": todo})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(result)
}

func TimeSpentEndPoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := Client.Database("gotodo").Collection("todos")

	var todo models.Todo
	err := collection.FindOne(context.TODO(), models.Todo{ID: id}).Decode(&todo)
	if err != nil {
		fmt.Println("Cannot find todo")
		return
	}

	var updatedTodo models.Todo
	json.NewDecoder(req.Body).Decode(&updatedTodo)

	fmt.Println("updatedTodo")

	todo.TimeSpent = append(todo.TimeSpent, updatedTodo.TimeSpent[0])
	todo.TotalTimeSpent = todo.TotalTimeSpent + updatedTodo.TimeSpent[0].Duration

	result, err := collection.UpdateOne(context.TODO(), models.Todo{ID: id}, bson.M{"$set": todo})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(res).Encode(result)
}

// func GetTokenHandler(res http.ResponseWriter, req *http.Request) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// }
