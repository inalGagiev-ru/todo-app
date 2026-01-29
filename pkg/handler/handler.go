package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/inalGagiev-ru/todo-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

type errorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.authMiddleware())
	{
		users := api.Group("/users")
		{
			users.GET("/profile", h.getProfile)
			users.PUT("/profile", h.updateProfile)
			users.DELETE("/profile", h.deleteAccount)
		}

		tasks := api.Group("/tasks")
		{
			tasks.POST("/", h.createTask)
			tasks.GET("/", h.getAllTasks)
			tasks.GET("/:id", h.getTaskByID)
			tasks.PUT("/:id", h.updateTask)
			tasks.DELETE("/:id", h.deleteTask)
		}

		categories := api.Group("/categories")
		{
			categories.POST("/", h.createCategory)
			categories.GET("/", h.getAllCategories)
		}

		tags := api.Group("/tags")
		{
			tags.POST("/", h.createTag)
			tags.GET("/", h.getAllTags)
		}
	}

	return router
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message, Code: statusCode})
}
