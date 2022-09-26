package factory

import (
	"github.com/ayatkyo/alterra-agmc/day-10/database"
	"github.com/ayatkyo/alterra-agmc/day-10/internal/repository"
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
