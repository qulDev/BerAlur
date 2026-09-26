# Lingkungan dan deployment

Selalu pakai base `compose.yaml` bersama **satu** override. Jangan gabungkan override local dengan prod karena port PostgreSQL lokal akan ikut terbuka.

| Lingkungan | Environment file            | Override             | Akses                                |
| ---------- | --------------------------- | -------------------- | ------------------------------------ |
| Local      | `.env`, `APP_ENV=local`     | `compose.local.yaml` | Web loopback 3100, DB loopback 55432 |
| Dev        | `.env.dev`, `APP_ENV=dev`   | `compose.dev.yaml`   | Hanya web loopback VPS               |
| Prod       | `.env.prod`, `APP_ENV=prod` | `compose.prod.yaml`  | Hanya web loopback VPS               |

`APP_ENV` menentukan nama proyek/volume Compose. Dev dan prod wajib memiliki password, database dan env file sendiri. Jika password mengandung karakter khusus, URL-encode password pada `DATABASE_URL` agar cocok dengan `POSTGRES_PASSWORD`. Jangan memakai credential contoh pada server.

Contoh urutan di VPS yang sudah disiapkan, dari root proyek:

```sh
docker compose --env-file .env.prod -f compose.yaml -f compose.prod.yaml config --quiet
docker compose --env-file .env.prod -f compose.yaml -f compose.prod.yaml build
docker compose --env-file .env.prod -f compose.yaml -f compose.prod.yaml up -d --wait db
docker compose --env-file .env.prod -f compose.yaml -f compose.prod.yaml --profile tools run --build --rm migrate up
docker compose --env-file .env.prod -f compose.yaml -f compose.prod.yaml up -d --wait
```

Sebelum menjalankan urutan ini pada data produksi: siapkan TLS/domain di reverse proxy host, secret, role database terbatas, backup/restore, dan rollback aplikasi. Deployment tidak dilakukan otomatis oleh CI. Docker image di sini dikunci pada major/minor, dependency aplikasi pada versi exact dan lockfile; pembaruan image tetap harus melalui verifikasi.

Halaman awal tidak menyimpan file; storage volume/R2 ditambahkan saat fitur upload ada. Redis/queue, metrics stack, replica aplikasi, dan pemisahan host DB mengikuti kebutuhan yang terukur. Database dapat dipindah dengan mengganti `DATABASE_URL` dan pengaturan jaringan/sertifikat.
