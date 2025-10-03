package wiring

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SetupRouter(registry *Registry) *gin.Engine {
	r := gin.Default()
	templatePath := filepath.Join("internal", "presentation", "views", "templates", "*.html")
	r.LoadHTMLGlob(templatePath)
	r.GET("/", registry.ViewHandler.Home)

	v1 := r.Group("/api/v1")

	// Category Routes
	categoryGroup := v1.Group("/categories")
	{
		categoryGroup.POST("", registry.CategoryHandler.CreateCategory)
		categoryGroup.GET("", registry.CategoryHandler.GetCategories)
		categoryGroup.GET("/:id", registry.CategoryHandler.GetCategoryByID)
		categoryGroup.DELETE("/:id", registry.CategoryHandler.DeleteCategory)
		categoryGroup.PUT("/:id", registry.CategoryHandler.UpdateCategory)
	}

	cashbookGroup := v1.Group("/cashbooks")
	{
		cashbookGroup.GET("", registry.CashbookHandler.GetCashbooks)
		cashbookGroup.POST("", registry.CashbookHandler.CreateCashbook)
		cashbookGroup.GET("/:id", registry.CashbookHandler.GetCashbookByID)
		cashbookGroup.DELETE("/:id", registry.CashbookHandler.DeleteCashbook)
		cashbookGroup.PUT("/:id", registry.CashbookHandler.UpdateCashbook)
	}

	return r
}
