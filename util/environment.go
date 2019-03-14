package util

import (
	"os"
)

type Environment int

const (
	Local Environment = iota
	Dev
	Stg
	Prod
)

type Locale string

func (e Environment) IsLocal() bool {
	return e == Local
}

func (e Environment) IsProd() bool {
	return e == Prod
}

var Env = func() Environment {
	switch os.Getenv("ENV") {
	case "dev":
		return Dev
	case "stg":
		return Stg
	case "prod":
		return Prod
	default:
		return Local
	}
}()

