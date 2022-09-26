package repository

import (
	"context"
	"errors"

	"github.com/ayatkyo/alterra-agmc/day-10/internal/models"
	"github.com/ayatkyo/alterra-agmc/day-10/pkg/utils"
)

type Book interface {
	FindAll(c context.Context) ([]models.Book, error)
	FindByID(c context.Context, ID uint) (*models.Book, error)
	Store(c context.Context, data models.Book) models.Book
	Update(c context.Context, ID uint, data models.Book) (*models.Book, error)
	Destroy(c context.Context, ID uint)
}

type book struct {
	DB     []models.Book
	LastID uint
}

func NewBookRepository() Book {
	books := []models.Book{
		{
			ID:     1,
			Title:  "Software Architecture with C++",
			Author: "Adrian Ostrowsmodels",
			Year:   2021,
		},
		{
			ID:     2,
			Title:  "Modern C++ Programming Cookbook - Second Edition",
			Author: "Marius Bancila",
			Year:   2020,
		},
		{
			ID:     3,
			Title:  "Flutter Apprentice - Third Edition",
			Author: "Vincent Ngo, Kevin D Moore and Michael Katz",
			Year:   2022,
		},
	}

	return &book{
		DB:     books,
		LastID: uint(len(books)),
	}
}

func (r *book) FindAll(c context.Context) ([]models.Book, error) {
	return r.DB, nil
}

func (r *book) FindByID(c context.Context, ID uint) (*models.Book, error) {
	book, ok := utils.Find(r.DB, func(item models.Book) bool {
		return item.ID == ID
	})

	if !ok {
		return nil, errors.New("book not found")
	}

	return &book, nil
}

func (r *book) Store(c context.Context, data models.Book) models.Book {
	newBook := models.Book{
		Title:  data.Title,
		Author: data.Author,
		Year:   data.Year,
	}

	// add new book
	newBook.ID = r.LastID + 1
	r.DB = append(r.DB, newBook)

	// increment lastID
	r.LastID++

	return newBook
}

func (r *book) Update(c context.Context, ID uint, data models.Book) (*models.Book, error) {
	bookIdx, ok := utils.FindIndex(r.DB, func(item models.Book) bool {
		return item.ID == ID
	})
	if !ok {
		return nil, errors.New("book not found")
	}

	// update
	book := &r.DB[bookIdx]
	book.Title = data.Title
	book.Author = data.Author
	book.Year = data.Year

	return book, nil
}

func (r *book) Destroy(c context.Context, ID uint) {
	utils.Delete(&r.DB, func(item models.Book) bool {
		return item.ID == ID
	})
}
