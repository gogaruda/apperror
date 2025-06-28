package apperror

import (
	"github.com/gin-gonic/gin"
	"os"
)

func HandleHTTPError(c *gin.Context, err error) {
	debug := os.Getenv("GIN_MODE") == "debug"
	res := ToHTTPError(err, debug)

	resp := gin.H{
		"code":    res.Status,
		"status":  "error",
		"message": res.UserMessage,
	}
	if res.Debug != "" {
		resp["debug"] = res.Debug
	}

	c.JSON(res.Status, resp)
}
