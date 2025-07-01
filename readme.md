# 🚀 MKP-THT – Backend Technical Test

Halo Tim Rekrutmen,

Berikut adalah hasil pengerjaan **Technical Test Backend** untuk PT. Mitra Kasih Perkasa (MKP).

## 📁 Struktur Project

```
.
├── cmd/                # Entry point (main.go)
├── config/             # Konfigurasi aplikasi
├── database/
│   ├── connections/    # File koneksi database
│   └── migrations/     # File migrasi SQL (Goose)
├── design-db/          # Desain struktur database & ERD
├── design-sistem/      # Diagram sistem/arsitektur
├── fundamental-test/   # Jawaban fundamental programming
├── internal/
│   ├── controllers/    # Handler untuk endpoint
│   ├── dto/            # Data Transfer Object (request & response)
│   ├── middlewares/    # Middleware (auth, dsb)
│   ├── models/         # Model domain dan enum
│   ├── repositories/   # Akses data & query DB
│   ├── routes/         # Inisialisasi routing
│   └── services/       # Business logic layer
├── utils/              # Utility/helper functions
└── 1_MKP.postman_collection.json # Postman Collection untuk test API
```

## 🔧 Cara Menjalankan Project

### 1. Clone Repository

```bash
git clone https://github.com/username/mkp-tht.git
cd mkp-tht
```

### 2. Setup Environment

- Pastikan **PostgreSQL** berjalan di `localhost`.
- Buat database baru, contoh: `mkp`.
- Jalankan migrasi menggunakan **Goose**:

```bash
goose -dir database/migrations postgres "postgres://username:password@localhost:5432/mkp?sslmode=disable" up
```

### 3. Jalankan Aplikasi

```bash
go run cmd/main.go
```

## 🧪 Testing API

### 📮 Postman Collection

Import file berikut ke dalam Postman:
`1_MKP.postman_collection.json`

## 📄 Penjelasan Folder

### 📌 Fundamental Test

- **Lokasi**: `fundamental-test/`
- Berisi jawaban tantangan logika dan dasar algoritma.

### 🧠 System Design

- **Lokasi**: `design-sistem/`
- Diagram arsitektur sistem E-Ticketing dan flow data.

### 🗂️ Database Design

- **Lokasi**: `design-db/`
- ERD, relasi antar entitas, dan struktur tabel SQL.

## 🛠️ Tools & Teknologi

- **Golang (Go)** – Backend implementation
- **PostgreSQL** – Relational database
- **Goose** – Database migration tool
- **Postman** – API testing

## ❗ Catatan Teknis

- Project menggunakan prinsip **clean architecture sederhana**.
- **Password tidak dikembalikan** dalam response.
- **Validasi dilakukan secara berlapis** (DTO + constraint SQL).
- **Middleware tersedia** untuk autentikasi (jika dibutuhkan).
- Desain sistem memungkinkan ekspansi untuk integrasi pembayaran dan gate scanner (RFID/NFC).

Terima kasih atas kesempatan yang diberikan 🙏