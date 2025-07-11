# apperror

Sistem standar untuk penanganan error di aplikasi Go (Gin-based), dengan dukungan:

✅ Kode error yang konsisten  
✅ Mapping otomatis ke HTTP status + pesan aman  
✅ Logging internal saat debug  
✅ Respons JSON yang bisa dikustom `status`-nya (tidak hanya "error")

---

## 🚀 Instalasi

```bash
go get github.com/gogaruda/apperror@v1.3.0
````

---

## 📦 API

### ✅ Membuat Error Standar

```go
err := apperror.New(apperror.CodeValidationError, "Nama wajib diisi", nil)
```

### ✅ Membuat Error dengan HTTP Status

```go
err := apperror.New(apperror.CodeValidationError, "Nama wajib diisi", nil, http.StatusBadRequest)
```

### ✅ Membuat Error dengan Status JSON Dinamis

```go
err := apperror.NewWithStatus("TOKEN_EXPIRED", "Token kadaluarsa", nil, 401, "expired")
```

### ✅ Atau chaining:

```go
err := apperror.New("TOKEN_EXPIRED", "Token kadaluarsa", nil, 401).
	WithResponseStatus("expired")
```

---

## ✅ Tangani Error Otomatis (Gin)

```go
func SomeHandler(c *gin.Context) {
	err := doSomething()
	if err != nil {
		apperror.HandleHTTPError(c, err)
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}
```

---

## ✅ Cek Jenis Error

```go
if apperror.Is(err, apperror.CodeUnauthorized) {
	// Redirect ke login
}
```

---

## 🔁 Contoh Output JSON

```json
{
  "code": 401,
  "status": "expired",
  "message": "Token kadaluarsa"
}
```

Atau default (tanpa `ResponseStatus`):

```json
{
  "code": 500,
  "status": "error",
  "message": "Terjadi kesalahan internal"
}
```

---

## 🔍 Cara Kerja `HandleHTTPError`

1. Cek apakah error adalah `InitError`
2. Cek apakah ada mapping `Code → HTTP status + User message`
3. Jika tidak ditemukan:

  * Pakai HTTP status dari error (jika ada)
  * Default ke 500
4. Status JSON:

  * Jika `ResponseStatus` di-set → dipakai
  * Jika tidak → fallback ke `"error"`
5. Pesan internal hanya ditampilkan di log jika `GIN_MODE=debug`

---

## ✅ Rekomendasi Pemakaian

Gunakan `apperror` untuk:

* Validasi input/form
* Kesalahan otentikasi dan otorisasi
* Konflik resource
* Error dari database atau layanan eksternal
* Penanganan logika bisnis yang bisa diharapkan
