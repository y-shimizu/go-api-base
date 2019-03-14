package lifecycle

import (
	"github.com/y-shimizu/go-api-base/domain/models"
)

type UserRepository interface {
	Find(id int64) (*models.User, error)
}
