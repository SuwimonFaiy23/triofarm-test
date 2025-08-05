package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/SuwimonFaiy23/triofarm-test/internal/model"
	"github.com/SuwimonFaiy23/triofarm-test/internal/service"
	"github.com/SuwimonFaiy23/triofarm-test/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemHandler interface {
	CreateItem(c *gin.Context)
	UpdateItem(c *gin.Context)
	DeleteItem(c *gin.Context)
	UpdateIndexMenu(c *gin.Context)
	GetItemList(c *gin.Context)
}

type itemHandler struct {
	itemService service.ItemService
}

func NewItemHandler(itemService service.ItemService) ItemHandler {
	return &itemHandler{itemService: itemService}
}

func (h *itemHandler) CreateItem(c *gin.Context) {
	var payload model.ItemRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request parameters",
			Error:   err.Error(),
		})
		return
	}

	missing := []string{}
	if payload.MenuID <= 0 {
		missing = append(missing, "menu_id")
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

	if err := h.itemService.CreateItem(payload); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.APIResponse{
				Code:    http.StatusNotFound,
				Message: "Data not found",
				Error:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.APIResponse{
		Code:    http.StatusCreated,
		Message: "Item created successfully",
	})
}

func (h *itemHandler) UpdateItem(c *gin.Context) {
	var payload model.ItemRequest
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

	if err := h.itemService.UpdateItem(payload); err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Code:    http.StatusOK,
		Message: "Item updated successfully",
	})
}

func (h *itemHandler) DeleteItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request parameters",
			Error:   err.Error(),
		})
		return
	}

	if err := h.itemService.DeleteItem(id); err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Code:    http.StatusOK,
		Message: "Item deleted successfully",
	})
}

func (h *itemHandler) UpdateIndexMenu(c *gin.Context) {
	var payload model.UpdateItemRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request parameters",
			Error:   err.Error(),
		})
		return
	}

	if payload.ItemID <= 0 {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Missing required fields: item_id",
		})
		return
	}

	if err := h.itemService.UpdateIndexItemMenu(payload); err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Code:    http.StatusOK,
		Message: "Item updated successfully",
	})
}

func (h *itemHandler) GetItemList(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request parameters",
			Error:   err.Error(),
		})
		return
	}

	result, err := h.itemService.GetItemListByMenuID(id)
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
