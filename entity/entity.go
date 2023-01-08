// Package entity domain-rich structures with associative behaviors (through identifiers), representing an entity/actor
// in a domain process.
package entity

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate          *validator.Validate
	validateSingleton sync.Once
)

func init() {
	validateSingleton.Do(func() {
		validate = validator.New()
	})
}
