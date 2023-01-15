package user

import (
	"github.com/maestre3d/coinlog/storage"
)

type Repository interface {
	storage.Repository[User]
}
