package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/SuwimonFaiy23/triofarm-test/internal/model"
	"github.com/SuwimonFaiy23/triofarm-test/internal/service"
	"github.com/SuwimonFaiy23/triofarm-test/pkg/response"
	"github.com/gin-gonic/gin"
)

type MenuHandler interface {
	CreateMenu(c *gin.Context)
	UpdateMenu(c *gin.Context)
	DeleteMenu(c *gin.Context)
	GetMenuList(c *gin.Context)
}

type menuHandler struct {
	menuService service.MenuService
}

func NewMenuHandler(menuService service.MenuService) MenuHandler {
	return &menuHandler{menuService: menuService}
}

func (h *menuHandler) CreateMenu(c *gin.Context) {
	var payload model.MenuRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request parameters",
			Error:   err.Error(),
		})
		return
	}

	if strings.TrimSpace(payload.Name) == "" {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Missing required fields: name",
		})
		return
	}
	if len(strings.TrimSpace(payload.Name)) > 255{
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid name: the value exceeds the maximum length of 255 characters",
		})
		return
	}

	if err := h.menuService.CreateMenu(payload); err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.APIResponse{
		Code:    http.StatusCreated,
		Message: "Menu created successfully",
	})
}

func (h *menuHandler) UpdateMenu(c *gin.Context) {
	var payload model.MenuRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request parameters",
			Error:   err.Error(),
		})
		return
	}

	missing := []string{}
	if payload.ID <= 0 {
		missing = append(missing, "id")
	}
	if strings.TrimSpace(payload.Name) == "" {
		missing = append(missing, "name")
	}
	if len(missing) > 0 {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Missing required fields: " + strings.Join(missing, ", "),
		})
		return
	}
	if len(strings.TrimSpace(payload.Name)) > 255{
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid name: the value exceeds the maximum length of 255 characters",
		})
		return
	}

	if err := h.menuService.UpdateMenu(payload); err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Code:    http.StatusOK,
		Message: "Menu updated successfully",
	})
}

func (h *menuHandler) DeleteMenu(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id <= 0{
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request parameters",
			Error:   err.Error(),
		})
		return
	}

	if err := h.menuService.DeleteMenu(id); err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Code:    http.StatusOK,
		Message: "Menu deleted successfully",
	})
}

func (h *menuHandler) GetMenuList(c *gin.Context) {

	result, err := h.menuService.GetMenuList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    result,
	})
}
