export default function Home() {
  return (
    <main>
      <header className="wordmark">
        <span aria-hidden="true">b.</span> BerAlur
      </header>
      <section className="intro" aria-labelledby="title">
        <p className="eyebrow">LANGKAH PERTAMA</p>
        <h1 id="title">
          Beralur, mulai
          <br />
          dari fondasi.
        </h1>
        <p className="description">
          Ruang untuk membangun alur bisnis yang lebih tertata. Kerangka awal
          sudah tersedia; fitur bisnis akan tumbuh dari sini.
        </p>
        <nav aria-label="Pemeriksaan layanan">
          <a className="primary" href="/health">
            Periksa API <span aria-hidden="true">↗</span>
          </a>
          <a href="/ready">
            Periksa database <span aria-hidden="true">↗</span>
          </a>
        </nav>
      </section>
      <aside className="note">
        <span className="note-label">FONDASI AWAL</span>
        <p>
          Halaman ini adalah titik awal pengembangan. Akun, pelanggan, stok, dan
          transaksi belum tersedia.
        </p>
      </aside>
      <footer>
        <span>Alur yang baik dimulai dari dasar yang jelas.</span>
        <span>01 / FONDASI</span>
      </footer>
    </main>
  );
}
