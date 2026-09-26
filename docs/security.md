# Keamanan dan tenant

Fondasi ini hanya mengekspos health/readiness dan halaman awal. Tabel organisasi belum mempunyai endpoint publik. UUID organisasi dan query lookup berdasarkan ID **tidak memberikan otorisasi**.

Sebelum menambahkan endpoint bisnis:

- Putuskan autentikasi, session dan membership; tentukan organisasi aktif dari identitas yang sudah diverifikasi.
- Semua operasi data tenant wajib scope `organization_id`, termasuk detail, update, delete, export, task background dan relasi FK. Scope tidak boleh dipercaya hanya karena dikirim client.
- Foreign key antar data tenant harus menjaga organisasi yang sama, misalnya composite key `(organization_id, id)` dan FK composite.
- Tulis integration test tenant A/B dan role yang tidak berhak. Tentukan 404/403 secara konsisten tanpa membocorkan keberadaan data tenant lain.
- Bila memakai cookie session: `HttpOnly`, `Secure`, `SameSite` dan perlindungan CSRF untuk mutasi. Tentukan rate limit pada login dan endpoint mahal.
- Catat perubahan uang/stok/izin di audit log dalam transaksi yang sama. Jangan log password, token, body request, query string sensitif, atau credential database.

Saat ini API memakai timeout, header request ID, respons error generik, structured log tanpa body/query string, dan user container non-root. Nginx membatasi body menjadi 1 MiB dan menambahkan header dasar. CORS tidak dibuka karena web dan API memakai satu origin.

Konfigurasi dev/prod hanya mempublikasikan Nginx ke loopback VPS. Pasang TLS pada reverse proxy tepi sebelum akses publik. Database eksternal perlu `sslmode=verify-full` dan CA yang benar; `sslmode=disable` di contoh hanya untuk jaringan Docker lokal. Role database aplikasi dan role migration perlu dipisah saat menyiapkan produksi.

Scaffold ini belum dinyatakan siap menyimpan data pelanggan produksi.
