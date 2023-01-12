package domain

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	Validate          *validator.Validate
	validateSingleton sync.Once
)

func init() {
	validateSingleton.Do(func() {
		Validate = validator.New()
	})
}
