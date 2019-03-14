package infrastructure

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(5*time.Minute, 30*time.Second)
