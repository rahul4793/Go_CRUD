package manager

import (
	"crud-app/models"
	"crud-app/mongoDatabase"
	"crud-app/request"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserManager struct {
	collection   *mongo.Collection
	mongodbQuery *mongoDatabase.User
}

func NewUserManager(col *mongo.Collection) *UserManager {
	return &UserManager{collection: col}
}

// usermgr :=new(UserManager)  add this

func (m *UserManager) CreateUser(req request.CreateUserRequest) (*models.User, error) {
	return m.mongodbQuery.InsertUser(req)
}

func (m *UserManager) GetAllUsers(page, limit int64) ([]models.User, error) {
	return m.mongodbQuery.GetAllUsers(page, limit)
}

func (m *UserManager) GetUserByID(id string) (*models.User, error) {
	return m.mongodbQuery.GetUserByID(id)
}

func (m *UserManager) UpdateUser(id string, req request.UpdateUserRequest) error {
	return mongoDatabase.UpdateUser(id, req)
}

func (m *UserManager) DeleteUser(id string) error {
	return mongoDatabase.SoftDeleteUser(id)
}
