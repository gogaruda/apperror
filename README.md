# apperror

## 🚀 Instalasi

```bash
go get github.com/gogaruda/apperror@v1.2.2
```

---

## Penggunaan

### Membuat Error Standar

```go
import "github.com/gogaruda/apperror"

err := apperror.New(apperror.CodeValidationError, "Nama wajib diisi", nil)
```

### ✅ Membuat Error Custom

```go
err := apperror.NewWithStatus("USER_BANNED", "Akun Anda diblokir", nil, 403)
```

### ✅ Tangani Error Otomatis (di handler Gin)

```go
import (
  "github.com/gin-gonic/gin"
  "github.com/gogaruda/apperror"
)

func SomeHandler(c *gin.Context) {
    err := doSomething()
    if err != nil {
        apperror.HandleHTTPError(c, err)
        return
    }

    c.JSON(200, gin.H{"status": "ok"})
}
```

### ✅ Cek Kode Error (opsional)

```go
if apperror.Is(err, apperror.CodeUnauthorized) {
    // arahkan ke login
}
```

---

## 🧠 Apa yang Dilakukan HandleHTTPError

* Mencari mapping dari `Code -> HTTP status + UserMessage`
* Jika ditemukan: kirim response dengan message aman
* Jika tidak ditemukan:

  * Gunakan HTTP status custom dari error (jika ada)
  * Jika tidak ada: fallback 500
* Pesan internal (`Message`) hanya dicetak ke log jika `GIN_MODE=debug`
* Client hanya menerima `UserMessage` yang aman

### Contoh Log (saat debug aktif):

```
[ERROR] CODE=VALIDATION_ERROR | MESSAGE=Nama wajib diisi | DETAIL=input kosong
```

### Contoh Response:

```json
{
  "code": 400,
  "status": "error",
  "message": "Input tidak valid"
}
```

---

## 📌 Rekomendasi

Gunakan `apperror` untuk:

* Validasi form/input
* Kesalahan otorisasi/autentikasi
* Error database (tx fail, no rows, dll)
* Penanganan error bisnis khusus
