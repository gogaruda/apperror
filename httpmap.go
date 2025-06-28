package apperror

import "net/http"

type HTTPErrorMap struct {
	Code        string
	Status      int
	UserMessage string
}

var defaultErrorMap = []HTTPErrorMap{
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
	{CodeDBNoRows, http.StatusNotFound, "Data tidak tersedia"},
	{CodeDBConstraint, http.StatusConflict, "Gagal menyimpan data: constraint"},
	{CodeDBTxFailed, http.StatusInternalServerError, "Transaksi database gagal"},
	{CodeDBConnFailed, http.StatusServiceUnavailable, "Koneksi database gagal"},
	{CodeDBError, http.StatusInternalServerError, "Kesalahan database"},
	{CodeDBPrepareError, http.StatusInternalServerError, "Gagal prepare query database"},
}
