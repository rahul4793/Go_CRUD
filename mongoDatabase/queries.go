package mongoDatabase

import (
	"context"
	"errors"
	"time"

	"crud-app/models"
	"crud-app/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultTimeout = 5 * time.Second

func withTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, defaultTimeout)
}

func InsertUser(ctx context.Context, col *mongo.Collection, req request.CreateUserRequest) (*models.User, error) {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	if count, err := col.CountDocuments(ctx, bson.M{"email": req.Email, "isDeleted": false}); err != nil {
		return nil, err
	} else if count > 0 {
		return nil, errors.New("email already exists")
	}

	user := &models.User{
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

func GetAllUsers(ctx context.Context, col *mongo.Collection, page, limit int64) ([]models.User, error) {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	skip := (page - 1) * limit
	opts := options.Find().SetSkip(skip).SetLimit(limit)

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

// func GetAllUsers(col *mongo.Collection, page, limit int64) ([]models.User, error) {
// 	skip := (page - 1) * limit
// 	opts := options.Find().SetSkip(skip).SetLimit(limit)

// 	cursor, err := col.Find(nil, bson.M{"isDeleted": false}, opts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(nil)

// 	var users []models.User
// 	for cursor.Next(nil) {
// 		var u models.User
// 		if err := cursor.Decode(&u); err != nil {
// 			return nil, err
// 		}
// 		users = append(users, u)
// 	}
// 	return users, nil
// }

func GetUserByID(ctx context.Context, col *mongo.Collection, id string) (*models.User, error) {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := col.FindOne(ctx, bson.M{"_id": objID, "isDeleted": false}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(ctx context.Context, col *mongo.Collection, id string, req request.UpdateUserRequest) error {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"name": req.Name, "email": req.Email, "age": req.Age}}
	_, err = col.UpdateOne(ctx, bson.M{"_id": objID, "isDeleted": false}, update)
	return err
}

func SoftDeleteUser(ctx context.Context, col *mongo.Collection, id string) error {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = col.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"isDeleted": true}})
	return err
}
