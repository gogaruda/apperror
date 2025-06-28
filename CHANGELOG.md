# Changelog

## [v1.2.1] - (2025-06-28)
### Perbaikan
- Pindahkan debug ke log(tidak di response JSON)

## [v1.2.0] - (2025-06-28)
### Perbaikan
- Tambahkan custom error

## [v1.0.3] - (2025-06-27)
### Perbaikan
- Ganti semua response c.JSON ke c.Abort untuk menghentikan semua proses setelahnya

## [v1.0.2] - 2025-06-26
### Perbaikan
- Perbaiki typo HandleHTTPError

### Penambahan
- Tambah pesan error

## [v1.0.1] - 2025-06-24
### Perbaikan
- Update module github.com/gogaruda/apperror

## [v1.0.0] - 2025-06-24
### Rilis pertama
- Struktur `InitError` untuk error wrapping
- Fungsi `apperror.New` dan `apperror.Is`
- Integrasi handler untuk Gin (`HandleHTTPError`)
- Daftar kode error umum (DB, auth, resource, internal)