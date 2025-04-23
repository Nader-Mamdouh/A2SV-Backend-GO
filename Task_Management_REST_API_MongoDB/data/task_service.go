package data

import (
	"TaskManagementRESTAPI_MongoDB/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = "Task_management"
var collectionname = "Task"

var TaskCollection *mongo.Collection

func ConnectDB() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017/Task_management")
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	TaskCollection = client.Database(db).Collection(collectionname)
}

func CreateTask(task models.Task) (*mongo.InsertOneResult, error) {
	return TaskCollection.InsertOne(context.TODO(), task)
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := TaskCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskByID(id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var task models.Task
	err = TaskCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	return &task, err
}

func UpdateTask(id string, task models.Task) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = TaskCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": task},
	)
	return err
}

func DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = TaskCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}
