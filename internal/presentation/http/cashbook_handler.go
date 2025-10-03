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

type CashbookHandler struct {
	service *application.CashbookService
}

func NewCashbookHandler(s *application.CashbookService) *CashbookHandler {
	return &CashbookHandler{service: s}
}

func (h *CashbookHandler) GetCashbooks(c *gin.Context) {
	cashbooks, err := h.service.GetCashbooks(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve cashbooks"})
		return
	}

	var res []dto.CashbookResponse
	for _, cashbook := range cashbooks {
		res = append(res, dto.CashbookResponse{
			ID:          cashbook.ID,
			Name:        cashbook.Name,
			Description: cashbook.Description,
			Amount:      cashbook.Amount,
		})
	}

	c.JSON(http.StatusOK, res)
}

func (h *CashbookHandler) GetCashbookByID(c *gin.Context) {
	idStr := c.Param("id")
	cashbookIdUint64, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cashbook ID format"})
		return
	}
	cashbookId := uint(cashbookIdUint64)
	cashbook, err := h.service.GetCashbookById(c.Request.Context(), cashbookId)

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "cashbook not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := dto.CashbookResponse{
		ID:          cashbook.ID,
		Name:        cashbook.Name,
		Description: cashbook.Description,
		Amount:      cashbook.Amount,
	}

	c.JSON(http.StatusOK, res)
}

func (h *CashbookHandler) DeleteCashbook(c *gin.Context) {
	idStr := c.Param("id")
	cashbookIdUint64, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cashbook ID format"})
		return
	}

	cashbookId := uint(cashbookIdUint64)
	err = h.service.DeleteCashbook(c.Request.Context(), cashbookId)

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "cashbook not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cashbook deleted successfully"})
}

func (h *CashbookHandler) CreateCashbook(c *gin.Context) {
	var req dto.CashbookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cashbook, err := h.service.CreateCashbook(c.Request.Context(), req.Name, req.Description, req.Amount)
	if err != nil {
		if errors.Is(err, application.ErrNegativeCashbookAmount) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create cashbook"})
		return
	}

	res := dto.CashbookResponse{
		ID:          cashbook.ID,
		Name:        cashbook.Name,
		Description: cashbook.Description,
		Amount:      cashbook.Amount,
	}
	c.JSON(http.StatusCreated, res)
}

func (h *CashbookHandler) UpdateCashbook(c *gin.Context) {
	var req dto.CashbookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	idStr := c.Param("id")
	cashbookIdUint64, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cashbook ID format"})
		return
	}

	cashbookId := uint(cashbookIdUint64)
	cashbook, err := h.service.UpdateCashbook(c.Request.Context(), req.Name, req.Description, req.Amount, cashbookId)

	if err != nil {

		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if errors.Is(err, application.ErrNegativeCashbookAmount) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create cashbook"})
		return
	}

	res := dto.CashbookResponse{
		ID:          cashbook.ID,
		Name:        cashbook.Name,
		Description: cashbook.Description,
		Amount:      cashbook.Amount,
	}
	c.JSON(http.StatusCreated, res)
}
