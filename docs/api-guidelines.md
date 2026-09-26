# Kontrak API

Endpoint bisnis menggunakan `/api/v1`, nama resource plural, JSON `camelCase`, dan UTC RFC3339 untuk timestamp. DB memakai `snake_case`.

## Respons

Single resource dikirim langsung:

```json
{ "id": "uuid", "name": "Toko Maju" }
```

Collection:

```json
{
  "items": [],
  "pagination": { "page": 1, "limit": 20, "total": 0 }
}
```

Create memakai `201` dan resource yang dibuat; delete memakai `204` tanpa body. Hindari `success`, salinan HTTP status, timestamp transport, dan message sukses di body.

Error:

```json
{ "error": { "code": "NOT_FOUND", "message": "Resource not found" } }
```

Validation error menggunakan `422` dan menambahkan `fields` hanya jika diperlukan:

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid request",
    "fields": { "name": "Name is required" }
  }
}
```

Request ID berada pada header `X-Request-ID`. API membuat ID sendiri; header ini bukan identitas pengguna/tenant. Jangan bocorkan SQL, stack trace, credential atau isi error dependency.

## Request bisnis mendatang

- Resource CRUD: `GET/POST /customers`, `GET/PATCH/DELETE /customers/{id}`.
- Command yang memiliki aturan bisnis: `POST /orders/{id}/confirm`, bukan PATCH status tanpa validasi transisi.
- Page pagination: `page=1&limit=20`; batas maksimum awal 100. Nilai tidak valid ditolak, tidak diteruskan ke SQL mentah.
- Sort: `sort=-createdAt`; filter dan nama kolom sort wajib allowlist.
- Request mutasi wajib membatasi body, menolak JSON rusak/field tidak dikenal, dan memvalidasi field. Belum ada endpoint mutasi pada scaffold ini.
- Jangan menerima `organizationId` dari payload sebagai bukti hak akses. Scope harus berasal dari membership terautentikasi.
- List hanya memuat field yang diperlukan; relasi detail diambil pada endpoint detail.

## Endpoint yang benar-benar tersedia

| Endpoint                         | Hasil                                                                     |
| -------------------------------- | ------------------------------------------------------------------------- |
| `GET /health`                    | `200 {"status":"ok"}`, tidak mengakses DB                                 |
| `GET /ready`                     | `200 {"status":"ready"}` saat DB dapat diping; selain itu `503 NOT_READY` |
| Route tidak dikenal              | `404 NOT_FOUND`                                                           |
| Method salah pada route yang ada | `405 METHOD_NOT_ALLOWED`                                                  |
| Panic sebelum respons ditulis    | `500 INTERNAL_ERROR`                                                      |

Readiness memeriksa konektivitas DB, **bukan versi migrasi**. Jalankan migrasi secara eksplisit sebelum mengaktifkan fitur yang bergantung pada tabel baru. Lihat [OpenAPI](openapi.yaml).
