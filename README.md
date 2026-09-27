# BerAlur

Fondasi SaaS: **Next.js + TypeScript**, **Go + chi**, **PostgreSQL + pgx/sqlc/goose**, **pnpm**, **Docker Compose**, dan **Nginx**.

Saat ini tersedia halaman awal, health/readiness API, error JSON minimal, konfigurasi terpusat, logging, migrasi organisasi dasar, serta pemeriksaan FE/BE. Auth, RBAC, isolasi tenant untuk data bisnis, customer, inventory, order, invoice, dan payment belum diimplementasikan.

## Perintah singkat dengan Make

`Makefile` ini memakai skrip pnpm dan konfigurasi Docker Compose yang sudah ada; tidak membuat atau menimpa `.env`, tidak menghapus volume database, dan tidak menjalankan deployment produksi.

Prasyaratnya adalah GNU Make (`make`), Node.js 24.15.0, pnpm 10.34.5, Go 1.27, serta Docker Desktop bila menjalankan database atau stack. Pada Windows, pasang GNU Make yang menyediakan perintah `make` di `PATH`; PowerShell tetap dapat dipakai untuk menjalankannya. Resep Make tidak memiliki deteksi platform atau pengelolaan proses sendiri, sehingga juga cocok untuk GNU Make di Linux.

```powershell
make                 # bantuan saja; tidak mengubah layanan atau data
make install
make db-up
make db-migrate
make dev-api
```

Di terminal kedua jalankan `make dev-web`. Untuk menjalankan seluruh stack Docker, gunakan `make local`, lihat log dengan `make logs`, lalu hentikan dengan `make down`. Target `down` menjalankan Compose `down` tanpa `-v`, sehingga volume PostgreSQL tetap tersimpan.

Pemeriksaan perubahan:

```powershell
make check           # test, lint, typecheck, dan format-check; berhenti saat gagal
make build
make db-generate
make format           # satu-satunya target yang mengubah format file
```

E2E tetap terpisah. Stack harus aktif di `http://127.0.0.1:3100` dan browser Playwright perlu dipasang terlebih dahulu dengan `pnpm exec playwright install chromium`; sesudah itu jalankan `make e2e`. Bila Chromium Playwright tidak dapat dijalankan secara lokal, ikuti alternatif Chrome pada [testing](docs/testing.md).

## Jalankan dengan Docker

Prasyarat: Docker Desktop dengan Linux containers dan Compose v2. Jalankan dari folder ini.

```powershell
Copy-Item .env.example .env
make local
make db-migrate
```

Jangan menimpa `.env` bila sudah berisi konfigurasi sendiri. `.env.example` hanya untuk lokal.

- Web: <http://localhost:3100>
- API process: <http://localhost:3100/health>
- API + database: <http://localhost:3100/ready>
- PostgreSQL lokal: `127.0.0.1:55432`

```powershell
make logs
make down
```

`down` mempertahankan volume database. Jangan menambahkan `-v` pada database yang ingin disimpan.

## Pengembangan native

Gunakan Node.js **24.15.0** (lihat `.nvmrc`), pnpm **10.34.5**, dan Go **1.27**. Versi dependensi tercatat dalam lockfile. Docker tetap dipakai untuk PostgreSQL.

Di Windows dengan NVM for Windows, pasang Node sekali bila belum tersedia, lalu pilih versi proyek ini pada setiap terminal baru:

```powershell
nvm install 24.15.0
nvm use 24.15.0
node --version # v24.15.0
pnpm --version # 10.34.5
go version # go1.27.x
```

Setelah runtime sesuai, jalankan dari root checkout saat ini. `pnpm install --frozen-lockfile` membangun ulang junction dependency untuk lokasi ini tanpa mengubah versi pada lockfile.

```powershell
if (-not (Test-Path .env)) { Copy-Item .env.example .env }
make install
make db-up
make db-migrate
$env:DATABASE_URL = 'postgres://beralur:beralur_local_only@127.0.0.1:55432/beralur?sslmode=disable'
make dev-api
```

Di terminal kedua: `make dev-web`, lalu buka <http://localhost:3000>. Next.js meneruskan request API ke `127.0.0.1:8080`. API membaca environment proses, bukan otomatis membaca `.env`.

## Pemeriksaan

```powershell
make check
make build
make db-generate
pnpm exec playwright install chromium
make e2e
```

E2E membutuhkan stack Docker yang sedang berjalan. Detail integration test dan bukti eksekusi ada di [testing](docs/testing.md).

## Panduan

- [Arsitektur dan batas fondasi](docs/architecture.md)
- [Kontrak API dan error](docs/api-guidelines.md)
- [Coding dan aturan domain](docs/coding-guidelines.md)
- [Keamanan dan multi-tenancy](docs/security.md)
- [Database, migration, backup, restore](docs/database.md)
- [Local, dev, prod dan deployment](docs/environments.md)
- [Rencana dan progres inisialisasi](docs/implementation-plan.md)

Tidak ada deployment otomatis atau lisensi distribusi yang dipilih pada tahap ini.
