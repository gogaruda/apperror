package apperror

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func HandleHTTPError(c *gin.Context, err error) {
	debugMode := os.Getenv("APP_DEBUG") == "true"

	var initErr *InitError
	if errors.As(err, &initErr) {
		// Logging internal
		if debugMode {
			log.Printf("[ERROR] CODE=%s | MESSAGE=%s | DETAIL=%+v", initErr.Code, initErr.Message, initErr.Err)
		}

		// Cari mapping kode error
		for _, mapping := range defaultErrorMap {
			if mapping.Code == initErr.Code {
				c.JSON(mapping.Status, gin.H{
					"code":    mapping.Status,
					"status":  "error",
					"message": mapping.UserMessage,
				})
				return
			}
		}

		// Tidak ada mapping → gunakan custom status atau fallback 500
		status := initErr.HTTPStatus
		if status == 0 {
			status = http.StatusInternalServerError
		}
		c.JSON(status, gin.H{
			"code":    status,
			"status":  "error",
			"message": initErr.Message,
		})
		return
	}

	// Bukan InitError → fallback umum
	if debugMode {
		log.Printf("[UNHANDLED ERROR]: %+v", err)
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"status":  "error",
		"message": "Terjadi kesalahan server",
	})
}
