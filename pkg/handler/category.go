package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inalGagiev-ru/todo-app/pkg/request"
)

func (h *Handler) createCategory(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input request.CreateCategoryInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.services.CreateCategory(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (h *Handler) getAllCategories(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	categories, err := h.services.GetAllCategories()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}
