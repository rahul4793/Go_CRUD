package mongoDatabase

import (
	"context"
	"errors"
	"time"

	"crud-app/models"
	"crud-app/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	mongo *models.User
}

// update and create in one
// upsert
func (user *User) InsertUser(req request.CreateUserRequest) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := user.mongo.GetUserCollection()

	if count, err := col.CountDocuments(ctx, bson.M{"email": req.Email, "isDeleted": false}); err != nil {
		return nil, err
	} else if count > 0 {
		return nil, errors.New("email already exists")
	}

	userModel := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		Age:       req.Age,
		IsDeleted: false,
	}

	res, err := col.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

// aggreagation
// facet
// total in pagniantion
func (Queryis *User) GetAllUsers(page, limit int64) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	skip := (page - 1) * limit
	opts := options.Find().SetSkip(skip).SetLimit(limit)
	col := models.GetUserCollection()

	cursor, err := col.Find(ctx, bson.M{"isDeleted": false}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var u models.User
		if err := cursor.Decode(&u); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (Queryis *User) GetUserByID(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.User
	col := models.GetUserCollection()

	if err := col.FindOne(ctx, bson.M{"_id": objID, "isDeleted": false}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id string, req request.UpdateUserRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	col := models.GetUserCollection()

	update := bson.M{"$set": bson.M{"name": req.Name, "email": req.Email, "age": req.Age}}
	_, err = col.UpdateOne(ctx, bson.M{"_id": objID, "isDeleted": false}, update)
	return err
}

func SoftDeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	col := models.GetUserCollection()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = col.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"isDeleted": true}})
	return err
}
