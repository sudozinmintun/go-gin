package views

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ViewHandler struct {
	// Note: This handler doesn't use the Application layer services,
	// as it's only serving static information, but it could if needed.
}

func NewViewHandler() *ViewHandler {
	return &ViewHandler{}
}

func (h *ViewHandler) Home(c *gin.Context) {
	data := gin.H{
		"Title":   "Home Page",
		"AppName": "Accounting",
		"Time":    time.Now().Format(time.RFC1123),
	}

	c.HTML(http.StatusOK, "home.html", data)
}
