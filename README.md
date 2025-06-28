# apperror

## Instalasi

```bash
go get github.com/gogaruda/apperror@latest
```

---

## Penggunaan

### Membuat Error Standar (dengan kode bawaan)

```go
import "github.com/gogaruda/apperror"

// Contoh: validasi input gagal
err := apperror.New(apperror.CodeValidationError, "Nama wajib diisi", nil)
```

### Membuat Error Custom (kode dan status bebas)

```go
// Error khusus bisnis: user banned, gunakan HTTP 403
err := apperror.New("USER_BANNED", "Akun Anda diblokir", nil, 403)
```

### Tangani Error secara otomatis (cukup satu baris)

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/gogaruda/apperror"
)

func SomeHandler(c *gin.Context) {
    err := someBusinessLogic()
    if err != nil {
        apperror.HandleHTTPError(c, err)
        return
    }

    c.JSON(200, gin.H{"status": "ok"})
}
```

> Fungsi `HandleHTTPError()` akan otomatis membaca `GIN_MODE=debug` dan mengatur status/message sesuai error.

### Cek Jenis Error

```go
if apperror.Is(err, apperror.CodeUnauthorized) {
    // lakukan redirect login atau token refresh
}
```

### Output JSON Standar

```json
{
  "code": 400,
  "status": "error",
  "message": "Input tidak valid",
  "debug": "Nama wajib diisi" // jika GIN_MODE=debug
}
```

---

## Catatan Penting

* Jika kode error tidak ditemukan di mapping bawaan (`defaultErrorMap`):

    * Akan menggunakan `initErr.HTTPStatus` jika tersedia
    * Jika tidak, default ke `500`
    * Pesan tetap berasal dari `initErr.Message`

---

## Direkomendasikan Untuk

* Gin REST API
* Middleware error handler global
* Penanganan validasi, otorisasi, dan DB error