package wiring

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(registry *Registry) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	// Category Routes
	categoryGroup := v1.Group("/categories")
	{
		categoryGroup.POST("", registry.CategoryHandler.CreateCategory)
		categoryGroup.GET("", registry.CategoryHandler.GetCategories)
		categoryGroup.GET("/:id", registry.CategoryHandler.GetCategoryByID)
		categoryGroup.DELETE("/:id", registry.CategoryHandler.DeleteCategory)
	}

	return r
}
