package repository

import (
	"errors"

	"github.com/cagnotteApp/Backend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoRepository) SaveUser(user *domain.User) (*domain.User, error) {
	client := r.client
	result, err := client.Database(r.dbName).Collection("Users").InsertOne(r.context, user)

	if err != nil {
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return user, nil
}

func (r *MongoRepository) FindUser(accAddr string) (*domain.User, error) {

	client := r.client
	user := domain.User{}
	cur, err := client.Database(r.dbName).Collection("Users").Find(r.context, bson.M{"accaddr": accAddr})
	if err != nil {
		return &user, err
	}
	returnedValues := []domain.User{}

	err = cur.All(r.context, &returnedValues)

	if err != nil {
		return &user, err
	}
	if len(returnedValues) == 0 {
		return nil, errors.New("user not found in the db")
	}
	return &returnedValues[0], nil
}
