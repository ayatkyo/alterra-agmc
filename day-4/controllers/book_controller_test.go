package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/ayatkyo/alterra-agmc/day-4/lib/database"
	"github.com/ayatkyo/alterra-agmc/day-4/middlewares"
	"github.com/ayatkyo/alterra-agmc/day-4/models"
	"github.com/ayatkyo/alterra-agmc/day-4/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBookGetAll(t *testing.T) {
	c, res, _ := utils.CreateTestContext(http.MethodGet, "/books", "")

	if assert.NoError(t, BookGetAll(c)) {
		assert.Equal(t, http.StatusOK, res.Code)

		resJSON := utils.ParseResponseJSON(res.Body.String())
		assert.Equal(t, "Get all books", resJSON["message"])
	}
}

func TestBookGetByID(t *testing.T) {
	c, res, _ := utils.CreateTestContext(http.MethodGet, "/books/:id", "")

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Get expected result
	book := database.BookDB[0]

	if assert.NoError(t, BookGetByID(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())
		userRes := resJSON["data"].(map[string]any)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "Get book by ID", resJSON["message"])
		assert.Equal(t, book.ID, uint(userRes["id"].(float64)))
		assert.Equal(t, book.Title, userRes["title"])
		assert.Equal(t, book.Author, userRes["author"])
		assert.Equal(t, book.Year, uint(userRes["year"].(float64)))
	}
}

func TestBookGetByIDWithInvalidParam(t *testing.T) {
	c, res, _ := utils.CreateTestContext(http.MethodGet, "/books/:id", "")

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("abc")

	if assert.NoError(t, BookGetByID(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Cannot parse book ID", resJSON["message"])
	}
}

func TestBookGetByIDNotFound(t *testing.T) {
	c, res, _ := utils.CreateTestContext(http.MethodGet, "/books/:id", "")

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("5")

	if assert.NoError(t, BookGetByID(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Book not found", resJSON["message"])
	}
}

func TestBookStore(t *testing.T) {
	// prepare book data
	bookTest := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		Year:   2022,
	}
	body, _ := json.Marshal(bookTest)

	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPost, "/books", string(body))

	// Set JWT token
	token, _ := middlewares.CreateToken(2)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(BookStore)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())
		bookRes := resJSON["data"].(map[string]any)

		t.Log(res.Body.String())

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "Book created", resJSON["message"])
		assert.NotEmpty(t, bookRes["id"])
		assert.Equal(t, bookTest.Title, bookRes["title"])
		assert.Equal(t, bookTest.Author, bookRes["author"])
		assert.Equal(t, bookTest.Year, uint(bookRes["year"].(float64)))
	}
}

func TestBookStoreValidation(t *testing.T) {
	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPost, "/books", "")

	// Set JWT token
	token, _ := middlewares.CreateToken(2)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(BookStore)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NotEqual(t, "Book created", resJSON["message"])
	}
}

func TestBookStoreNoAuth(t *testing.T) {
	// prepare book data
	bookTest := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		Year:   2022,
	}
	body, _ := json.Marshal(bookTest)

	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPost, "/books", string(body))

	if assert.NoError(t, middlewares.JWTMiddleware(BookStore)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Missing or malformed JWT", resJSON["message"])
	}
}

func TestBookUpdate(t *testing.T) {
	// prepare book data
	bookTest := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		Year:   2022,
	}
	body, _ := json.Marshal(bookTest)

	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPut, "/books/:id", string(body))

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("2")

	// Set JWT token
	token, _ := middlewares.CreateToken(2)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(BookUpdate)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())
		bookRes := resJSON["data"].(map[string]any)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "Book updated", resJSON["message"])
		assert.NotEmpty(t, bookRes["id"])
		assert.Equal(t, bookTest.Title, bookRes["title"])
		assert.Equal(t, bookTest.Author, bookRes["author"])
		assert.Equal(t, bookTest.Year, uint(bookRes["year"].(float64)))
	}
}

func TestBookUpdateValidation(t *testing.T) {
	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPut, "/books/:id", "")

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Set JWT token
	token, _ := middlewares.CreateToken(1)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(BookUpdate)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NotEqual(t, "Book created", resJSON["message"])
	}
}

func TestBookUpdateNoAuth(t *testing.T) {
	// prepare book data
	bookTest := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		Year:   2022,
	}
	body, _ := json.Marshal(bookTest)

	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPut, "/books/:id", string(body))

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("2")

	if assert.NoError(t, middlewares.JWTMiddleware(BookUpdate)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Missing or malformed JWT", resJSON["message"])
	}
}

func TestBookDestroy(t *testing.T) {
	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodDelete, "/books/:id", "")

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Set JWT token
	token, _ := middlewares.CreateToken(1)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(BookDestroy)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, fmt.Sprintf("Successfully delete book with ID %d", 1), resJSON["message"])
	}
}
func TestBookDestroyInvalidID(t *testing.T) {
	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodDelete, "/books/:id", "")

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("abc")

	// Set JWT token
	token, _ := middlewares.CreateToken(1)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(BookDestroy)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Cannot parse book ID", resJSON["message"])
	}
}

func TestBookDestroyNoAuth(t *testing.T) {
	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodDelete, "/books/:id", "")

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("2")

	if assert.NoError(t, middlewares.JWTMiddleware(BookDestroy)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Missing or malformed JWT", resJSON["message"])
	}
}
