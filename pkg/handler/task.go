package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/inalGagiev-ru/todo-app/pkg/repository"
	"github.com/inalGagiev-ru/todo-app/pkg/request"
)

func (h *Handler) createTask(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input request.CreateTaskInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := (*h.services).CreateTask(userID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *Handler) getAllTasks(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	filters := repository.TaskFilters{
		Status: c.Query("status"),
	}

	if catID := c.Query("category_id"); catID != "" {
		if catIDInt, err := strconv.Atoi(catID); err == nil {
			categoryID := uint(catIDInt)
			filters.CategoryID = &categoryID
		}
	}

	if tagIDsStr := c.Query("tag_ids"); tagIDsStr != "" {
		var tagIDs []uint
		tagIDStrs := strings.Split(tagIDsStr, ",")
		for _, tagIDStr := range tagIDStrs {
			if tagIDInt, err := strconv.Atoi(tagIDStr); err == nil {
				tagIDs = append(tagIDs, uint(tagIDInt))
			}
		}
		filters.TagIDs = tagIDs
	}

	tasks, err := h.services.GetAllTasks(userID, filters)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) getTaskByID(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	task, err := (*h.services).GetTaskByID(userID, uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) updateTask(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input request.UpdateTaskInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := (*h.services).UpdateTask(userID, uint(id), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) deleteTask(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = (*h.services).DeleteTask(userID, uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Task deleted successfully",
	})
}
