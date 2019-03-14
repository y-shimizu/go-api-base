package api

import (
	"github.com/y-shimizu/go-api-base/domain/lifecycle"
	repo "github.com/y-shimizu/go-api-base/infrastructure/lifecycle"
)

var (
	userRepository lifecycle.UserRepository = repo.NewUserRepositoryMySQL()
)
