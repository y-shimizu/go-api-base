package lifecycle

import (
	"github.com/y-shimizu/go-api-base/domain/models"
)

type UserRepositoryMySQL struct{}

func NewUserRepositoryMySQL() UserRepositoryMySQL {
	return UserRepositoryMySQL{}
}

func (this UserRepositoryMySQL) Find(id int64) (*models.User, error) {
	//var user *models.User
	//record, err := mysql.NewTUserDao().FindByID(id.UserID)
	//TODO: record to model
	//user = recordToEntity(record)
	return nil, nil
}

