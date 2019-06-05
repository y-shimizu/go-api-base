package lifecycle

import (
	"github.com/y-shimizu/go-api-base/domain/models"
	"github.com/y-shimizu/go-api-base/infrastructure/mysql"
)

type UserRepositoryMySQL struct{}

func NewUserRepositoryMySQL() UserRepositoryMySQL {
	return UserRepositoryMySQL{}
}

func (this UserRepositoryMySQL) Find(id int64) (*models.User, error) {
	var user *models.User
	record, err := mysql.NewUserDao().FindById(id)
	//TODO: record to model
	user = recordToEntity(record)
	return user, err
}

func recordToEntity(record *mysql.User) *models.User {
	user := models.User{
		Id: record.ID,
		Name: record.Name,
		Email: record.Email,
	}
	return &user
}