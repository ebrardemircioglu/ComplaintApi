package internal

import (
	"ComplaintSystem/UserService/internal/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(e *echo.Echo, service *Service) {
	handler := &Handler{service: service}

	g := e.Group("/user")
	g.GET("/:id", handler.GetByID)
	g.POST("/", handler.Create)
	g.GET("/", handler.GetAll)
	g.PUT("/:id", handler.Update)
	g.PATCH("/:id", handler.PartialUpdate)
	g.DELETE("/:id", handler.Delete)
}

// statusbadgateway
func (h *Handler) Create(c echo.Context) error {
	var userRequestModel *types.UserRequestModel
	err := c.Bind(&userRequestModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Can not bind JSON: " + err.Error(),
		})
	}
	id, err := h.service.Create(c.Request().Context(), userRequestModel)

	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Can not create user: " + err.Error(),
		})
	}
	response := map[string]interface{}{
		"createdId": id,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *Handler) GetByID(c echo.Context) error {

	id := c.Param("id")
	result, err := h.service.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, result)
}
func (h *Handler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := h.service.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}
func (h *Handler) Update(c echo.Context) error {
	id := c.Param("id")
	var user types.UserUpdateModel
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	err = h.service.Update(c.Request().Context(), id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User updated successfully",
	})
}
func (h *Handler) PartialUpdate(c echo.Context) error {
	id := c.Param("id")
	var user types.UserUpdateModel
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	err = h.service.PartialUpdate(c.Request().Context(), id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Update failed"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Contact information updated"})
}

/*
	func (h *Handler) GetAll(c echo.Context) error {
		name := c.QueryParam("name")
		surname := c.QueryParam("surname")
		address := c.QueryParam("address")
		result, err := h.service.GetAll(c.Request().Context(), name, surname, address)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, result)
	}
*/
func (h *Handler) GetAll(c echo.Context) error {
	var user types.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	result, err := h.service.GetAll(c.Request().Context(), user.Name, user.Surname, user.Address)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
