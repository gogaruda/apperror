package apperror

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type HTTPErrorMap struct {
	Code        string
	Status      int
	UserMessage string
}

var defaultErrorMap = []HTTPErrorMap{
	{CodeUserNotFound, http.StatusNotFound, "User tidak ditemukan"},
	{CodeResourceNotFound, http.StatusNotFound, "Data tidak ditemukan"},
	{CodeInvalidInput, http.StatusBadRequest, "Input tidak valid"},
	{CodeValidationError, http.StatusBadRequest, "Validasi gagal"},
	{CodeBadRequest, http.StatusBadRequest, "Permintaan tidak sesuai"},
	{CodeUserConflict, http.StatusConflict, "User sudah ada"},
	{CodeResourceConflict, http.StatusConflict, "Data bentrok atau duplikat"},
	{CodeUnauthorized, http.StatusUnauthorized, "Anda harus login"},
	{CodeForbidden, http.StatusForbidden, "Akses ditolak"},
	{CodeNotImplemented, http.StatusNotImplemented, "Fitur belum tersedia"},
	{CodeTimeout, http.StatusGatewayTimeout, "Permintaan melebihi waktu tunggu"},
	{CodeDependencyError, http.StatusBadGateway, "Kesalahan dari layanan eksternal"},
	{CodeDBNoRows, http.StatusNotFound, "Data tidak tersedia"},
	{CodeDBConstraint, http.StatusConflict, "Gagal menyimpan data: constraint"},
	{CodeDBTxFailed, http.StatusInternalServerError, "Transaksi database gagal"},
	{CodeDBConnFailed, http.StatusServiceUnavailable, "Koneksi database gagal"},
	{CodeDBError, http.StatusInternalServerError, "Kesalahan database"},
	{CodeInternalError, http.StatusInternalServerError, "Terjadi kesalahan internal"},
	{CodeRoleNotFound, http.StatusBadRequest, "Role tidak ditemukan"},
	{CodeAuthNotFound, http.StatusNotFound, "Username/email atau password salah"},
}

var debugMode = os.Getenv("GIN_MODE") == "debug"

func HandleHTTError(c *gin.Context, err error) {
	var initErr *InitError
	if errors.As(err, &initErr) {
		// Logging lengkap
		log.Printf("[ERROR] %s | MESSAGE: %s | DETAIL: %+v\n", initErr.Code)

		// Temukan mapping HTTP response
		for _, mapping := range defaultErrorMap {
			if mapping.Code == initErr.Code {
				resp := gin.H{
					"code":    mapping.Status,
					"status":  "error",
					"message": mapping.UserMessage,
				}

				if debugMode {
					resp["debug"] = initErr.Message
				}

				c.JSON(mapping.Status, resp)
				return
			}
		}

		log.Printf("[UNMAPPED ERROR CODE] %s | MESSAGE: %s", initErr.Code, initErr.Message)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Terjadi kesalahan server",
			"debug":   initErr.Message,
		})
	}

	// Unknown error (bukan initError)
	log.Printf("[UNHANDLED ERROR]: %+v", err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  "error",
		"message": "Terjadi kesalahan server",
		"debug":   err.Error(),
	})
}
