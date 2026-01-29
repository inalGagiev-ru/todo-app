package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/inalGagiev-ru/todo-app/pkg/utils"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func (e errorResponse) Error() string {
	return fmt.Sprintf("%s (code: %d)", e.Message, e.Code)
}

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
			return
		}

		if len(headerParts[1]) == 0 {
			newErrorResponse(c, http.StatusUnauthorized, "token is empty")
			return
		}

		userID, err := utils.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		c.Set(userCtx, userID)
		c.Next()
	}
}

func getUserId(c *gin.Context) (uint, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errorResponse{Message: "user id not found", Code: http.StatusInternalServerError}
	}

	userID, ok := id.(uint)
	if !ok {
		return 0, errorResponse{Message: "user id is of invalid type", Code: http.StatusInternalServerError}
	}

	return userID, nil
}
