package user

import (
	"github.com/maestre3d/coinlog/domain"
)

type Repository interface {
	domain.Repository[User]
}
