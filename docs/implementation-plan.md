# Inisialisasi BerAlur

**Tujuan:** fondasi lokal yang berjalan, dapat diperiksa ulang, dan mengikuti [arsitektur](architecture.md) yang telah dibahas.

**Pelaksanaan:** langsung di proyek baru `beralur/`, di luar `sources/`. Tidak ada publikasi atau deployment ke server.

- [x] API: tulis test untuk konfigurasi tidak valid, health tanpa DB, readiness gagal, error 404/405/panic dan request ID; jalankan sebelum implementasi. Tambahkan config, router dan entrypoint server; jalankan Go test/vet/build.
- [x] Database: buat migrasi identitas organisasi minimal dan query lookup berdasarkan ID; generate sqlc. Jalankan migrasi up/down/up pada DB khusus pemeriksaan dan uji organisasi lain tidak terambil oleh ID yang berbeda.
- [x] Web: siapkan Next.js, ESLint, Prettier, TypeScript dan Vitest/RTL. Uji tautan pemeriksaan API sebelum menambah halaman awal. Jalankan test, lint, typecheck dan build.
- [x] Operasional: siapkan Compose local/dev/prod, Nginx, Dockerfile, CI, panduan kontrak/security/testing/database dan smoke test Playwright. Jalankan stack lengkap dan tes melalui proxy.

Perhatian pemeriksaan: readiness tidak boleh sukses ketika DB mati; pesan error tidak membocorkan detail DB; database dan port aplikasi tidak dipublikasikan oleh konfigurasi prod; query tenant selalu membutuhkan scope; perintah restore tidak diarahkan ke database kerja.

Review independen menemukan DNS Nginx yang perlu diperbarui ketika container diganti dan petunjuk native yang belum menyalin `.env`. Keduanya telah diperbaiki. Hasil verifikasi serta batas buktinya tercatat di [testing](testing.md).
