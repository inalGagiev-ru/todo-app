package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inalGagiev-ru/todo-app/pkg/request"
)

func (h *Handler) signUp(c *gin.Context) {
	var input request.SignUpInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.SignUp(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) signIn(c *gin.Context) {
	var input request.SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.SignIn(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
