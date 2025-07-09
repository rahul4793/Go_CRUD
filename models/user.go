// package models

// import (
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type User struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
// 	Name      string             `bson:"name" json:"name"`
// 	Email     string             `bson:"email" json:"email"`
// 	Age       int                `bson:"age" json:"age"`
// 	IsDeleted bool               `bson:"isDeleted" json:"isDeleted"`
// }

//	func GetUserCollection() *mongo.Collection {
//		return database.MongoDatabase.Collection("users")
//	}
package models

import (
	"crud-app/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Age       int                `bson:"age" json:"age"`
	IsDeleted bool               `bson:"isDeleted" json:"isDeleted"`
}

func GetUserCollection() *mongo.Collection {
	if database.MongoDatabase == nil {
		panic("MongoDatabase not initialized. Call database.Connect() first.")
	}
	return database.MongoDatabase.Collection("users")
}
