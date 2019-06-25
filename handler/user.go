package handler

import (
	"echo-swag/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// AddUser godoc
// @Summary Add a user
// @Description add user from JSON POST request
// @Tags users
// @Accept  json
// @Produce  json
// @Param User body model.AddUser true "Add User"
// @Success 200 {object} model.User
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/add [post]
func (h *Handler) AddUser(c echo.Context) error {
	var addUser model.AddUser
	if err := c.Bind(&addUser); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := addUser.Validation(); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	user := model.User{
		Name:  addUser.Name,
		Email: addUser.Email,
	}
	lastID, err := user.Insert()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	user.ID = lastID
	return c.JSON(http.StatusOK, user)
}

// GetUser godoc
// @Summary Get a user
// @Description get user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/{id} [get]
func (h *Handler) GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	users, err := model.UserOne(aid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, users)
}

// ListAllUsers godoc
// @Summary List all users
// @Description list all users
// @Tags users
// @Produce  json
// @Success 200 {array} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/all [get]
func (h *Handler) ListAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, model.UsersAll())
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update by json user
// @Tags users
// @Accept  json
// @Produce  json
// @Param  id path int true "User ID"
// @Param  user body model.UpdateUser true "Update user"
// @Success 200 {object} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/{id} [patch]
func (h *Handler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var updateUser model.UpdateUser
	if err := c.Bind(&updateUser); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err)
	}
	user := model.User{
		ID:    aid,
		Name:  updateUser.Name,
		Email: updateUser.Email,
	}
	err = user.Update()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.String(http.StatusOK, "User "+id+" updated")
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete by user ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param  id path int true "user ID" Format(int64)
// @Success 204 {object} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = model.Delete(aid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	res := map[string]string{
		"id":     id,
		"status": "deleted",
	}
	return c.JSON(http.StatusOK, res)
}
