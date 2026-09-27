# Testing

## Unit dan build

`pnpm test`, `pnpm lint`, `pnpm typecheck`, `pnpm build`, dan `pnpm format:check` dijalankan dari root dengan Node 24.15.0, pnpm 10.34.5 dan Go 1.27. Pada Windows dengan NVM for Windows, jalankan `nvm use 24.15.0` sebelum memasang dependency atau memeriksa proyek.

- Go `testing` + Testify: konfigurasi, pemisahan health/readiness, error 404/405/500, request ID, dan DB unavailable.
- Vitest + React Testing Library: tautan pemeriksaan menggunakan origin yang sama.
- Playwright: halaman, navigasi, health, readiness dan kontrak error melalui Nginx.

Unit test readiness menyimulasikan kegagalan ping. Itu bukan bukti akses PostgreSQL nyata; integration test terpisah di bawah memeriksanya.

## PostgreSQL nyata

Siapkan database khusus yang telah dimigrasi; **jangan gunakan data produksi atau database kerja**. Contoh berikut membuat database pemeriksaan unik di PostgreSQL Compose, menguji urutan migration up/down/up, lalu menghapus database pemeriksaan itu saja.

```powershell
$checkDatabase = "beralur_check_$PID"
docker compose -f compose.yaml -f compose.local.yaml up -d --wait db
try {
  docker compose -f compose.yaml -f compose.local.yaml exec -T db sh -c 'createdb --username="$POSTGRES_USER" "$1"' -- $checkDatabase
  $env:DATABASE_URL = "postgres://beralur:beralur_local_only@db:5432/$checkDatabase?sslmode=disable"
  docker compose -f compose.yaml -f compose.local.yaml --profile tools run --build --rm migrate up
  docker compose -f compose.yaml -f compose.local.yaml --profile tools run --build --rm migrate down
  docker compose -f compose.yaml -f compose.local.yaml --profile tools run --build --rm migrate up
  $env:TEST_DATABASE_URL = "postgres://beralur:beralur_local_only@127.0.0.1:55432/$checkDatabase?sslmode=disable"
  Push-Location apps/api
  try { go test -tags integration ./... } finally { Pop-Location }
} finally {
  docker compose -f compose.yaml -f compose.local.yaml exec -T db sh -c 'dropdb --if-exists --username="$POSTGRES_USER" "$1"' -- $checkDatabase
}
```

Test membuat dua organisasi dalam transaksi yang di-rollback, memeriksa lookup berdasarkan ID, ID yang tidak ada, dan constraint nama kosong. Ini belum membuktikan otorisasi/isolasi tenant pada endpoint bisnis karena endpoint dan autentikasinya belum ada.

Go race detector dijalankan di CI Linux. Untuk pemeriksaan lokal Windows, race detector memerlukan C compiler; unit test biasa tetap bisa dijalankan tanpa itu.

## E2E

```powershell
pnpm exec playwright install chromium
pnpm test:e2e
```

Bila browser unduhan tidak dapat dijalankan di mesin lokal, gunakan Chrome yang sudah terpasang: `$env:PLAYWRIGHT_CHANNEL = 'chrome'`, lalu `pnpm test:e2e`. CI tetap memakai Chromium dari Playwright.

Stack harus aktif di `http://127.0.0.1:3100`; ubah `BASE_URL` bila port berbeda. Untuk pemeriksaan kegagalan dependency, hentikan hanya service `db` proyek ini, verifikasi `/ready` menjadi 503 sementara `/health` tetap 200, lalu hidupkan DB kembali.

`./tests/integration/proxy-replacement.ps1` memeriksa Nginx mengikuti pergantian alamat IP API pada network/container sementara, tanpa mengubah database atau stack lokal. Image `beralur-local-api` harus sudah dibangun.

CI menyertakan lint/typecheck/unit/build, generate sqlc, PostgreSQL integration, build container dan E2E. Hasil CI remote baru terbukti setelah workflow dijalankan di repository tujuan; hasil lokal tidak dianggap hasil CI remote.

## Hasil inisialisasi — 26 September 2026

Berhasil dijalankan lokal:

- Unit Go dan Vitest/RTL; lint, typecheck, Go vet, build API dan production build Next.js.
- Generate sqlc dan integration test menggunakan PostgreSQL 18 pada database `beralur_test`.
- Migrasi up/down/up pada database pemeriksaan; backup dan restore ke database baru. Penolakan restore ke database aplikasi maupun database yang sudah ada juga diperiksa.
- Validasi ketiga override Compose; build image API, web, migration; empat service lokal healthy.
- Database dihentikan: `/health` tetap 200, `/ready` menjadi 503; setelah DB kembali, `/ready` pulih ke 200.
- Test pergantian IP API awalnya gagal, kemudian lulus setelah Nginx memakai DNS resolver Docker.
- Satu E2E Playwright lulus melalui Chrome yang terpasang (`PLAYWRIGHT_CHANNEL=chrome`); tampilan desktop dan mobile diperiksa. Download kelengkapan Chromium sempat gagal dan executable unduhannya tidak dapat diluncurkan (`spawn UNKNOWN`), sehingga browser bawaan Playwright belum terverifikasi pada host ini.
- File `.env`, dependency, build output, backup, dan hasil test terabaikan oleh Git.

Go race detector dan workflow GitHub masih menunggu eksekusi CI Linux. Tidak ada bukti deployment VPS, load test, auth/RBAC atau isolasi endpoint bisnis pada tahap ini.
