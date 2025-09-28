package http

import (
	"errors"
	"net/http"
	"strconv"

	"example.com/accounting/internal/application"
	"example.com/accounting/internal/domain"
	"example.com/accounting/internal/presentation/dto"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *application.CategoryService
}

func NewCategoryHandler(s *application.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.service.CreateCategory(c.Request.Context(), req.Name, req.Description, req.Unit)
	if err != nil {
		if errors.Is(err, application.ErrInvalidCategoryName) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if errors.Is(err, application.ErrCategoryAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create category"})
		return
	}

	res := dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Unit:        category.Unit,
	}

	c.JSON(http.StatusCreated, res)
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	categories, err := h.service.GetCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve categories"})
		return
	}

	var res []dto.CategoryResponse
	for _, cat := range categories {
		res = append(res, dto.CategoryResponse{
			ID:          cat.ID,
			Name:        cat.Name,
			Description: cat.Description,
			Unit:        cat.Unit,
		})
	}
	c.JSON(http.StatusOK, res)
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	categoryIdUint64, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID format"})
		return
	}

	category, err := h.service.GetCategoryByID(c.Request.Context(), uint(categoryIdUint64))

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Unit:        category.Unit,
	}

	c.JSON(http.StatusOK, res)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	categoryIdUint64, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID format"})
		return
	}
	categoryId := uint(categoryIdUint64)
	err = h.service.DeleteCategory(c.Request.Context(), categoryId)

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	categoryIdUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID format"})
		return
	}
	categoryId := uint(categoryIdUint64)

	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category, err := h.service.UpdateCategory(c.Request.Context(), req.Name, req.Description, req.Unit, categoryId)

	if err != nil {
		// FIX 1: Check for the domain.ErrNotFound from the Service/Repository
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // Returns "category not found"
			return
		}

		if errors.Is(err, application.ErrCategoryAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		// Check for application errors
		if errors.Is(err, application.ErrInvalidCategoryName) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Generic error handler
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update category"}) // Changed message to update
		return
	}

	res := dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Unit:        category.Unit,
	}
	c.JSON(http.StatusOK, res)
}
