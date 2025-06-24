# apperror

Modul `apperror` adalah utilitas standar untuk menangani error di aplikasi backend berbasis Go, khususnya dalam arsitektur web (REST API) menggunakan Gin. Modul ini menyediakan standar penanganan error internal, pelaporan ke client, dan logging.

---

## 🔧 Fitur

- Standardisasi kode error (terstruktur)
- Support log internal dan pesan aman untuk client
- Debug mode untuk development
- Integrasi mudah dengan Gin
- Dukungan error wrapping & unwrapping (`errors.As`, `apperror.Is`)

---

## 📦 Instalasi

```bash
go get github.com/gogaruda/apperror
````
---

## 🚀 Cara Penggunaan

### 1. Buat error di repository layer

```go
return apperror.New(
    apperror.CodeDBError,
    fmt.Sprintf("gagal query role_id %v", roleID),
    err,
)
```

### 2. Tangani di handler controller

```go
data, err := s.GetByID(id)
if err != nil {
    apperror.HandleHTTPError(c, err)
    return
}
```

### 3. (Opsional) Cek tipe error dengan `Is`

```go
if apperror.Is(err, apperror.CodeUserNotFound) {
    // bisa return 404, atau log khusus
}
```

---

## 🐞 Debug Mode

Untuk menampilkan pesan internal (`debug`) di response JSON saat development, aktifkan env var:

```bash
APP_DEBUG=true
```

---

## 🔐 Best Practice

* Jangan tampilkan `err.Error()` langsung ke user.
* Gunakan `apperror.New` untuk membungkus semua error.
* Gunakan `apperror.Is` saat ingin handle khusus berdasarkan kode error.

---

## 📤 Contoh Response (Production)

```json
{
  "code": 404,
  "status": "error",
  "message": "User tidak ditemukan"
}
```

## 🧪 Contoh Response (Debug Mode)

```json
{
  "code": 500,
  "status": "error",
  "message": "Kesalahan database",
  "debug": "gagal query role_id 123"
}
```
