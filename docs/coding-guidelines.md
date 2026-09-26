# Coding dan aturan domain

1. Gunakan fungsi langsung dan standard library sebelum menambah abstraksi atau dependency.
2. Go handler menangani HTTP; aturan bisnis tinggal dekat domain. Jangan menambah interface jika belum ada kebutuhan substitusi nyata.
3. Query milik tenant harus menyertakan `organization_id`; baca aturan [security](security.md) sebelum fitur pertama.
4. Perubahan schema wajib migration. Edit query sumber lalu jalankan `pnpm db:generate`; jangan edit generated Go.
5. Jangan gunakan floating point untuk uang. Representasi integer/satuan mata uang atau NUMERIC diputuskan bersama desain invoice/payment, termasuk aturan pembulatan.
6. Konfirmasi order, pengurangan stok, pencatatan pembayaran, dan audit terkait wajib atomik ketika fitur tersebut ditambahkan.
7. Test perilaku kritis: izin, scope tenant, uang, stok, transisi status, dan kegagalan transaksi. Hindari target coverage tanpa alasan.
8. Unit test berada dekat source. Test lintas aplikasi berada di `tests/e2e`. Tidak perlu folder kosong untuk fitur yang belum dibuat.
9. `gofmt`, `go vet`, ESLint, Prettier, TypeScript dan build harus bersih. CI menjadi gerbang awal; Git hook tidak diwajibkan.
10. Kontrak publik tetap minimal dan kompatibel. Tambahkan keputusan arsitektur baru hanya saat ada tradeoff nyata.

Login/session, membership, RBAC, lifecycle order, stok negatif, nomor invoice, pembatalan dan pembayaran parsial belum diputuskan. Jangan menganggap halaman awal atau UUID organisasi sebagai implementasi aturan tersebut.
