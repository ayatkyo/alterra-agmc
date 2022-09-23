package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/ayatkyo/alterra-agcm/day-4/config"
	"github.com/ayatkyo/alterra-agcm/day-4/middlewares"
	"github.com/ayatkyo/alterra-agcm/day-4/models"
	"github.com/ayatkyo/alterra-agcm/day-4/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserGetAll(t *testing.T) {
	c, res, _ := utils.CreateTestContext(http.MethodGet, "/users", "")

	if assert.NoError(t, UserGetAll(c)) {
		assert.Equal(t, http.StatusOK, res.Code)

		resJSON := utils.ParseResponseJSON(res.Body.String())
		assert.Equal(t, "Get all users", resJSON["message"])
	}
}

func TestUserGetByID(t *testing.T) {
	c, res, _ := utils.CreateTestContext(http.MethodGet, "/users/:id", "")

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Get expected result
	var user models.User
	config.DB.Where("id = ?", 1).First(&user)

	if assert.NoError(t, UserGetByID(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())
		userRes := resJSON["data"].(map[string]any)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "Get user by ID", resJSON["message"])
		assert.Equal(t, user.ID, uint(userRes["id"].(float64)))
		assert.Equal(t, user.Fullname, userRes["fullname"])
		assert.Equal(t, user.Username, userRes["username"])
		assert.Equal(t, user.Email, userRes["email"])
	}
}
func TestUserGetByIDWithInvalidParam(t *testing.T) {
	c, res, _ := utils.CreateTestContext(http.MethodGet, "/users/:id", "")

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("abc")

	if assert.NoError(t, UserGetByID(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Cannot parse user ID", resJSON["message"])
	}
}

func TestUserStore(t *testing.T) {
	// Prepare test user
	pass := "testuserpassword"
	userTest := models.User{
		Fullname: "User Test",
		Username: "user.test",
		Email:    "user.test@gmail.com",
		Password: utils.BcryptMustMake(pass),
	}
	body, _ := json.Marshal(userTest)

	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPost, "/users", string(body))

	if assert.NoError(t, UserStore(c)) {
		// Get expected result
		var user models.User
		config.DB.Where("username = ?", userTest.Username).First(&user)

		resJSON := utils.ParseResponseJSON(res.Body.String())
		userRes := resJSON["data"].(map[string]any)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "User created", resJSON["message"])
		assert.NotEmpty(t, userRes["id"])
		assert.Equal(t, user.Fullname, userRes["fullname"])
		assert.Equal(t, user.Username, userRes["username"])
		assert.Equal(t, user.Email, userRes["email"])
	}
}

func TestUserStoreValidation(t *testing.T) {
	c, res, _ := utils.CreateTestContext(http.MethodPost, "/users", "")

	if assert.NoError(t, UserStore(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NotEqual(t, "User created", resJSON["message"])
	}
}

func TestUserStoreExistsUsername(t *testing.T) {
	// Prepare test user
	pass := "testuserpassword"
	userTest := models.User{
		Fullname: "User Test",
		Username: "ayat",
		Email:    "user.test@gmail.com",
		Password: utils.BcryptMustMake(pass),
	}
	body, _ := json.Marshal(userTest)

	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPost, "/users", string(body))

	if assert.NoError(t, UserStore(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Username or email already exists", resJSON["message"])
	}
}

func TestUserStoreExistsEmail(t *testing.T) {
	// Prepare test user
	pass := "testuserpassword"
	userTest := models.User{
		Fullname: "User Test",
		Username: "user.test",
		Email:    "mawad@gmail.com",
		Password: utils.BcryptMustMake(pass),
	}
	body, _ := json.Marshal(userTest)

	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPost, "/users", string(body))

	if assert.NoError(t, UserStore(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Username or email already exists", resJSON["message"])
	}
}

func TestUserUpdate(t *testing.T) {
	// Prepare test user
	userTest := models.User{
		Fullname: "User Test",
		Username: "alfath.test",
		Email:    "alfath@gmail.com",
	}
	body, _ := json.Marshal(userTest)

	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPut, "/users/:id", string(body))

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("2")

	// Set JWT token
	token, _ := middlewares.CreateToken(2)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(UserUpdate)(c)) {
		var user models.User
		config.DB.Where("id = ?", 2).First(&user)

		resJSON := utils.ParseResponseJSON(res.Body.String())
		userRes := resJSON["data"].(map[string]any)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "Your data has been updated", resJSON["message"])
		assert.NotEmpty(t, userRes["id"])
		assert.Equal(t, userTest.Fullname, userRes["fullname"])
		assert.Equal(t, userTest.Username, userRes["username"])
		assert.Equal(t, userTest.Email, userRes["email"])
	}
}

func TestUserUpdateInvalidParam(t *testing.T) {
	// Prepare test user
	userTest := models.User{
		Fullname: "User Test",
		Username: "user.test",
		Email:    "alfath@gmail.com",
	}
	body, _ := json.Marshal(userTest)

	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPut, "/users/:id", string(body))

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("abc")

	if assert.NoError(t, UserUpdate(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Cannot parse user ID", resJSON["message"])
	}
}

func TestUserUpadeValidation(t *testing.T) {
	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPut, "/users/:id", "")

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("2")

	// Set JWT token
	token, _ := middlewares.CreateToken(1)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(UserUpdate)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NotEqual(t, "Your data has been updated", resJSON["message"])
	}
}

func TestUserUpdateOtherUser(t *testing.T) {
	// Prepare test user
	userTest := models.User{
		Fullname: "User Test",
		Username: "alfath.test",
		Email:    "alfath@gmail.com",
	}
	body, _ := json.Marshal(userTest)

	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPut, "/users/:id", string(body))

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("2")

	// Set JWT token
	token, _ := middlewares.CreateToken(1)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(UserUpdate)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Cannot edit other user data", resJSON["message"])
	}
}

func TestUserUpdateUniqueUsername(t *testing.T) {
	// Prepare test user
	userTest := models.User{
		Fullname: "User Test",
		Username: "mawad",
		Email:    "alfath@gmail.com",
	}
	body, _ := json.Marshal(userTest)

	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPut, "/users/:id", string(body))

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("2")

	// Set JWT token
	token, _ := middlewares.CreateToken(2)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(UserUpdate)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Username or email already exists", resJSON["message"])
	}
}

func TestUserUpdateUniqueEmail(t *testing.T) {
	// Prepare test user
	userTest := models.User{
		Fullname: "User Test",
		Username: "alfath",
		Email:    "ayat@gmail.com",
	}
	body, _ := json.Marshal(userTest)

	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodPut, "/users/:id", string(body))

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("2")

	// Set JWT token
	token, _ := middlewares.CreateToken(2)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(UserUpdate)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Username or email already exists", resJSON["message"])
	}
}

func TestUserDestroy(t *testing.T) {
	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodDelete, "/users/:id", "")

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Set JWT token
	token, _ := middlewares.CreateToken(1)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(UserDestroy)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())

		var total int64
		config.DB.Model(&models.User{}).Where("id = ?", 1).Count(&total)
		assert.Less(t, int(total), 1)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "Your account has been deleted", resJSON["message"])
	}
}

func TestUserDestroyInvalidParam(t *testing.T) {
	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodDelete, "/users/:id", "")

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("abc")

	// Set JWT token
	token, _ := middlewares.CreateToken(1)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(UserDestroy)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Cannot parse user ID", resJSON["message"])
	}
}

func TestUserDestroyOtherUser(t *testing.T) {
	// Create context
	c, res, _ := utils.CreateTestContext(http.MethodDelete, "/users/:id", "")

	// Set param
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Set JWT token
	token, _ := middlewares.CreateToken(2)
	c.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))

	if assert.NoError(t, middlewares.JWTMiddleware(UserDestroy)(c)) {
		resJSON := utils.ParseResponseJSON(res.Body.String())
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "Cannot delete other user data", resJSON["message"])
	}
}
