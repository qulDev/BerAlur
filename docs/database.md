# Database

Go menerima koneksi hanya lewat `DATABASE_URL`. Host `db` digunakan di Compose; pengembangan native menggunakan `127.0.0.1:55432`. PostgreSQL 18 memakai volume `/var/lib/postgresql`.

Migrasi pertama hanya menyediakan identitas organisasi (`id`, `name`, `created_at`). UUID dibuat oleh PostgreSQL. Tabel bisnis, membership, audit dan aturan lifecycle belum dibuat.

```powershell
pnpm db:migrate
pnpm db:generate
```

SQL sumber: `apps/api/db/queries`; migrasi goose: `apps/api/db/migrations`; output sqlc: `apps/api/db/generated`. sqlc membaca sisi Up dari migrasi goose. Commit query, migrasi dan output generator bersama-sama.

Migrasi tidak dijalankan oleh setiap replica API. Pada deployment, jalankan satu job migration dengan credential yang sesuai sebelum aplikasi memakai skema baru. Migrasi yang sudah dibagikan tidak diubah; buat migrasi lanjutan. Rollback `00001` menghapus tabel organisasi: gunakan hanya pada database pemeriksaan kosong. Produksi mengutamakan forward-fix dan backup teruji.

## Backup dan latihan restore lokal

Jalankan dari root proyek dengan PowerShell:

```powershell
./scripts/backup-db.ps1
./scripts/restore-db.ps1 -BackupFile ./backups/FILE.dump -RestoreDatabase beralur_restore_check
```

Script restore selalu membuat **database baru**, menolak nama database aplikasi dan gagal bila database tujuan sudah ada. Tidak ada `--clean` atau penimpaan database kerja. Periksa hasil dengan `psql` sebelum backup dianggap terbukti dapat dipulihkan. Backup mengandung data: simpan di tempat privat; folder `backups/` diabaikan Git.

Belum ada demo seed karena belum ada model user/customer/product. Tambahkan seed lokal idempotent bersamaan dengan domain pertama, bukan data pelanggan produksi.
