package repository

import (
	"github.com/maestre3d/coinlog/entity"
)

type User interface {
	Repository[entity.User]
}
