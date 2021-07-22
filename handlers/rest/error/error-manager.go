package resterror

import (
	"hexa-go/infra/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendHTTPError Manage http response error
func SendHTTPError(c *gin.Context, err *error.HexagoError) {
	var status int

	switch err.Kind {
	case "Validation":
		status = http.StatusBadRequest
		c.JSON(status, gin.H{"error": err.String()})
	case "NotFound":
		c.Status(http.StatusNotFound)
	default:
		status = http.StatusInternalServerError
		c.JSON(status, gin.H{"error": err.String()})
	}
}
