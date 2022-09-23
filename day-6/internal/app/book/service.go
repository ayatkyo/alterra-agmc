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
	FindByID(c context.Context, ID uint) (*models.Book, error)
	Store(c context.Context, data *models.Book) (models.Book, error)
	Update(c context.Context, ID uint, data *models.Book) (*models.Book, error)
	Destroy(c context.Context, ID uint) error
}

func NewSevice(f *factory.Factory) Service {
	return &service{
		BookRepository: f.BookRepository,
	}
}

func (s *service) FindAll(c context.Context) ([]models.Book, error) {
	books, err := s.BookRepository.FindAll(c)
	return books, err
}

func (s *service) FindByID(c context.Context, ID uint) (*models.Book, error) {
	book, err := s.BookRepository.FindByID(c, ID)
	return book, err
}

func (s *service) Store(c context.Context, data *models.Book) (models.Book, error) {
	book := s.BookRepository.Store(c, *data)
	return book, nil
}

func (s *service) Update(c context.Context, ID uint, data *models.Book) (*models.Book, error) {
	book, err := s.BookRepository.Update(c, ID, *data)
	return book, err
}

func (s *service) Destroy(c context.Context, ID uint) error {
	book, err := s.BookRepository.FindByID(c, ID)
	if err != nil {
		return err
	}

	s.BookRepository.Destroy(c, book.ID)
	return nil
}
