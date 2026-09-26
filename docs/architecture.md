# Fondasi BerAlur

Keputusan ini merangkum percakapan **Analisis Pasar SaaS Pertama** dan permintaan untuk menginisialisasi proyek, 26 September 2026.

## Cakupan awal

Monorepo pnpm berisi Next.js + TypeScript (`apps/web`) dan modular monolith Go + chi (`apps/api`). PostgreSQL diakses lewat pgx dan query sqlc; perubahan skema menggunakan goose. Docker Compose menjalankan PostgreSQL, API, web, dan Nginx. API hanya menerima alamat database dari `DATABASE_URL`.

Satu tabel `organizations` menjadi identitas tenant minimal: UUID, nama, dan waktu dibuat. Belum ada endpoint organisasi atau data bisnis yang dapat diakses publik. Skema customer, inventory, order, invoice, payment, autentikasi, membership, RBAC, dan audit bisnis didesain dalam pekerjaan berikutnya. Fondasi ini belum menyediakan isolasi tenant yang lengkap.

Browser → Nginx → Next.js atau Go API → PostgreSQL. Saat pengembangan native, Next.js meneruskan `/api/*`, `/health`, dan `/ready` ke API sehingga browser tetap memakai satu origin.

## Batas modul

- Modul Go ditambahkan di `internal/<domain>` ketika ada fitur nyata. HTTP handler mengurai input dan memanggil aturan domain; query SQL berada di `db/queries`.
- Frontend memakai App Router. Komponen dan logic suatu fitur tinggal bersama fitur tersebut ketika sudah dibutuhkan.
- Tidak ada generic repository, base service, response wrapper universal, Redis, queue, microservice, atau folder domain kosong.
- File hasil sqlc berada di `apps/api/db/generated` dan tidak diedit manual.

## Kriteria selesai

Web dapat dibangun; API mempunyai `/health` dan `/ready`, request ID, JSON logging, timeout, graceful shutdown, serta error minimal. Unit test FE/BE, lint, typecheck, build, migrasi dan query pada PostgreSQL nyata, serta smoke test lewat Nginx dapat dijalankan ulang. CI menjalankan pemeriksaan yang sama tanpa deployment otomatis.

## Rujukan implementasi

- [Next.js installation](https://nextjs.org/docs/app/getting-started/installation)
- [Next.js standalone output](https://nextjs.org/docs/app/api-reference/config/next-config-js/output)
- [Next.js Vitest](https://nextjs.org/docs/app/guides/testing/vitest)
- [sqlc generation](https://docs.sqlc.dev/en/latest/howto/generate.html)
- [Compose startup order](https://docs.docker.com/compose/how-tos/startup-order/)
