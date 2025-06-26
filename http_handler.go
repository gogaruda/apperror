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
	// General
	{CodeInternalError, http.StatusInternalServerError, "Terjadi kesalahan internal"},
	{CodeInvalidInput, http.StatusBadRequest, "Input tidak valid"},
	{CodeBadRequest, http.StatusBadRequest, "Permintaan tidak sesuai"},
	{CodeValidationError, http.StatusBadRequest, "Validasi gagal"},
	{CodeUnauthorized, http.StatusUnauthorized, "Anda harus login"},
	{CodeForbidden, http.StatusForbidden, "Akses ditolak"},
	{CodeNotImplemented, http.StatusNotImplemented, "Fitur belum tersedia"},
	{CodeTimeout, http.StatusGatewayTimeout, "Permintaan melebihi waktu tunggu"},
	{CodeDependencyError, http.StatusBadGateway, "Kesalahan dari layanan eksternal"},
	{CodeEncodingError, http.StatusInternalServerError, "Gagal encoding data"},
	{CodeDecodingError, http.StatusBadRequest, "Gagal decoding data"},
	{CodeParseError, http.StatusBadRequest, "Gagal parsing data"},
	{CodeBindError, http.StatusBadRequest, "Gagal binding data"},
	{CodeMarshalError, http.StatusInternalServerError, "Gagal konversi data"},
	{CodeUnmarshalError, http.StatusBadRequest, "Gagal membaca data"},
	{CodePrepareError, http.StatusInternalServerError, "Gagal mempersiapkan data"},

	// Resource
	{CodeUserNotFound, http.StatusNotFound, "User tidak ditemukan"},
	{CodeUserConflict, http.StatusConflict, "User sudah terdaftar"},
	{CodeUsernameConflict, http.StatusConflict, "Username sudah terdaftar"},
	{CodeEmailConflict, http.StatusConflict, "Email sudah terdaftar"},
	{CodeResourceNotFound, http.StatusNotFound, "Data tidak ditemukan"},
	{CodeResourceConflict, http.StatusConflict, "Data bentrok atau duplikat"},
	{CodeRoleNotFound, http.StatusBadRequest, "Role tidak ditemukan"},
	{CodeAuthNotFound, http.StatusNotFound, "Username/email atau password salah"},
	{CodeTokenInvalid, http.StatusUnauthorized, "Token tidak valid"},
	{CodeTokenExpired, http.StatusUnauthorized, "Token telah kedaluwarsa"},
	{CodePermissionDenied, http.StatusForbidden, "Tidak memiliki izin"},
	{CodeInvalidCredential, http.StatusUnauthorized, "Kredensial tidak valid"},

	// Database
	{CodeDBNoRows, http.StatusNotFound, "Data tidak tersedia"},
	{CodeDBConstraint, http.StatusConflict, "Gagal menyimpan data: constraint"},
	{CodeDBTxFailed, http.StatusInternalServerError, "Transaksi database gagal"},
	{CodeDBConnFailed, http.StatusServiceUnavailable, "Koneksi database gagal"},
	{CodeDBError, http.StatusInternalServerError, "Kesalahan database"},
	{CodeDBPrepareError, http.StatusInternalServerError, "Gagal prepare query database"},
}

var debugMode = os.Getenv("APP_DEBUG") == "true"

func HandleHTTPError(c *gin.Context, err error) {
	var initErr *InitError
	if errors.As(err, &initErr) {
		// Logging lengkap
		log.Printf("[ERROR] %s | MESSAGE: %s | DETAIL: %+v\n", initErr.Code, initErr.Message, initErr.Err)

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
