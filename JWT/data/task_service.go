package data

import (
	"context"
	"JWT/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTasks(collection *mongo.Collection) ([]models.Task, error){

		cursor, err := collection.Find(context.TODO(), bson.M{})
		if err != nil{
			return nil, err
		}
		defer cursor.Close(context.TODO())

		var tasks []models.Task
		if err = cursor.All(context.TODO(), &tasks); err != nil{
			return nil, err
		}

		if tasks == nil{
			tasks = []models.Task{}
		}

		return tasks, nil
	
}

func GetTask(collection *mongo.Collection, taskID string) (models.Task, error){
	var task models.Task
	err := collection.FindOne(context.TODO(), bson.M{"id": taskID}).Decode(&task)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			return task, err
		}
	}

	return task, nil
}

func UpdateTask(collection *mongo.Collection, taskID string, update bson.M) (int64, int64, error){
	result, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"id": taskID},
		bson.M{"$set": update},
	)

	if err != nil{
		return 0, 0, err
	}
	return result.MatchedCount, result.ModifiedCount, nil
	
}


func RemoveTask(collection *mongo.Collection, taskID string) (int64, error){
	
	result, err := collection.DeleteOne(context.TODO(), bson.M{"id": taskID})
	if err != nil{
		return 0, err
	}

	return result.DeletedCount, nil

}

func AddTask(collection *mongo.Collection, newTask models.Task) (any, error) {
	result, err := collection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}