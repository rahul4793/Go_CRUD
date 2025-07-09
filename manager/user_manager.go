package manager

import (
	"context"

	"crud-app/models"
	"crud-app/mongoDatabase"
	"crud-app/request"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserManager struct {
	collection *mongo.Collection
}

func NewUserManager(col *mongo.Collection) *UserManager {
	return &UserManager{collection: col}
}

// usermgr :=new(UserManager)  add this

func (m *UserManager) CreateUser(ctx context.Context, req request.CreateUserRequest) (*models.User, error) {
	return mongoDatabase.InsertUser(ctx, m.collection, req)
}

// func (m *UserManager) GetAllUsers(ctx context.Context, page, limit int64) ([]models.User, error) {
// 	return mongoDatabase.GetAllUsers(ctx, m.collection, page, limit)
// }

func (m *UserManager) GetAllUsers(page, limit int64) ([]models.User, error) {
	return mongoDatabase.GetAllUsers(m.collection, page, limit)
}

func (m *UserManager) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return mongoDatabase.GetUserByID(ctx, m.collection, id)
}

func (m *UserManager) UpdateUser(ctx context.Context, id string, req request.UpdateUserRequest) error {
	return mongoDatabase.UpdateUser(ctx, m.collection, id, req)
}

func (m *UserManager) DeleteUser(ctx context.Context, id string) error {
	return mongoDatabase.SoftDeleteUser(ctx, m.collection, id)
}
