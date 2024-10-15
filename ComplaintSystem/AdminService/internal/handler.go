package internal

import (
	"ComplaintSystem/AdminService/internal/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(e *echo.Echo, service *Service) {
	handler := &Handler{service: service}

	g := e.Group("/admin")
	g.POST("/", handler.Create)
	g.GET("/:companyName", handler.GetByCompanyName)
	g.PUT("/:id", handler.Update)
	g.PATCH("/:companyName", handler.PartialUpdate)
	g.GET("/", handler.GetAll)
	g.DELETE("/:id", handler.Delete)
}
func (h *Handler) Create(c echo.Context) error {
	var admin *types.AdminRequestModel
	err := c.Bind(&admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Can not bind JSON: " + err.Error(),
		})
	}
	id, err := h.service.Create(c.Request().Context(), admin)

	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Can not create admin: " + err.Error(),
		})
	}
	response := map[string]interface{}{
		"createdId": id,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *Handler) GetAll(c echo.Context) error {
	admin, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, admin)
}
func (h *Handler) GetByCompanyName(c echo.Context) error {

	companyName := c.Param("companyName")
	//result, err := h.service.GetByCompanyName(c.Request().Context(), companyName)
	admin, err := h.service.GetByCompanyName(c.Request().Context(), companyName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if admin == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Company name not found"})
	}
	adminResponse := ToAdminResponse(admin)
	return c.JSON(http.StatusOK, adminResponse)

	//return c.JSON(http.StatusOK, result)
}
func (h *Handler) Update(c echo.Context) error {
	companyName := c.Param("companyName")
	var admin types.AdminUpdateModel
	err := c.Bind(&admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	err = h.service.Update(c.Request().Context(), companyName, admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Admin updated successfully",
	})
}

func (h *Handler) PartialUpdate(c echo.Context) error {
	companyName := c.Param("companyName")
	var admin types.AdminUpdateModel
	err := c.Bind(&admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	err = h.service.PartialUpdate(c.Request().Context(), companyName, admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Update failed"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": " employee updated"})
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
		"message": "Admin deleted successfully",
	})
}
