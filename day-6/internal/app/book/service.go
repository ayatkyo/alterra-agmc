package book

import (
	"context"

	"github.com/ayatkyo/alterra-agcm/day-6/internal/factory"
	"github.com/ayatkyo/alterra-agcm/day-6/internal/models"
	"github.com/ayatkyo/alterra-agcm/day-6/internal/repository"
)

type service struct {
	BookRepository repository.Book
}

type Service interface {
	FindAll(c context.Context) ([]models.Book, error)
}

func NewSevice(f *factory.Factory) Service {
	return &service{
		BookRepository: f.BookRepository,
	}
}

func (s *service) FindAll(c context.Context) ([]models.Book, error) {
	users, err := s.BookRepository.FindAll(c)
	return users, err
}
