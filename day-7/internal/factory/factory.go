package factory

import (
	"github.com/ayatkyo/alterra-agcm/day-7/database"
	"github.com/ayatkyo/alterra-agcm/day-7/internal/repository"
)

type Factory struct {
	UserRepository repository.User
	BookRepository repository.Book
}

func NewFactory() *Factory {
	db := database.DB

	return &Factory{
		UserRepository: repository.NewUserRepository(db),
		BookRepository: repository.NewBookRepository(),
	}
}
