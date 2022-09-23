package factory

import (
	"github.com/ayatkyo/alterra-agcm/day-6/database"
	"github.com/ayatkyo/alterra-agcm/day-6/internal/repository"
)

type Factory struct {
	UserRepository repository.User
}

func NewFactory() *Factory {
	db := database.DB

	return &Factory{
		UserRepository: repository.NewUserRepository(db),
	}
}
