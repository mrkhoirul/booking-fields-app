# Booking Lapangan API

[![Go Version](https://img.shields.io/badge/Go-1.25.3-00ADD8?logo=go&logoColor=white)](https://golang.org/)
[![Fiber Version](https://img.shields.io/badge/Fiber-v3-blue?logo=fiber)](https://gofiber.io/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18-blue?logo=postgresql)](https://hub.docker.com/_/postgres)
[![Docker Compose](https://img.shields.io/badge/Docker--Compose-Ready-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)

RESTful API untuk sistem **booking lapangan olahraga** berbasis **Golang (Fiber v3)** dan **PostgreSQL**.  
Proyek ini dikembangkan sebagai **Take Home Test – Backend Developer** di PT. Sagara Asia Teknologi.

## Fitur Utama

- **Authentication (JWT)**
  - Register & Login user
  - Role: `user` dan `admin`
- **Manajemen Lapangan (Fields)**
  - CRUD data lapangan (admin only)
- **Booking Lapangan**
  - Validasi agar tidak overlapping
  - Status awal booking: `pending`
- **Mock Payment (Fake Endpoint)**
  - Endpoint untuk ubah status booking jadi `paid`

## Tech Stack

| Komponen | Teknologi |
|-----------|------------|
| Bahasa Pemrograman | Go 1.25.3 |
| Framework | Fiber v3 |
| Database | PostgreSQL 18 |
| ORM | GORM |
| Authentication | JWT |
| Deployment | Docker & Docker Compose |
| API Test | Postman Collection |

## Instalasi & Menjalankan Proyek

### Prasyarat
- Sudah terinstal **Docker** & **Docker Compose**
- Port `8080` (app) dan `5432` (PostgreSQL) tidak sedang digunakan

### Jalankan Menggunakan Docker

1. Clone repository:
   ```bash
   git clone https://github.com/mrkhoirul/booking-fields-app.git
   cd booking-fields-app
   ```

2. Jalankan aplikasi:
   ```bash
   docker compose up --build
   ```

3. Server akan berjalan di:
   ```
   http://localhost:8080
   ```

4. Jalankan di background (opsional):
   ```bash
   docker compose up -d
   ```

5. Hentikan container:
   ```bash
   docker compose down
   ```

## Environment Variables

| Variable | Deskripsi | Contoh |
|-----------|------------|--------|
| `APP_PORT` | Port aplikasi | `8080` |
| `DB_HOST` | Host PostgreSQL | `localhost` |
| `DB_PORT` | Port PostgreSQL | `5432` |
| `DB_USER` | Username PostgreSQL | `postgres` |
| `DB_PASSWORD` | Password PostgreSQL | `postgres` |
| `DB_NAME` | Nama database | `bookingdb` |
| `JWT_SECRET` | Secret key JWT | `secret_key_jwt` |
| `JWT_EXPIRE_HOURS` | Durasi token dalam jam | `72` |
| `ADMIN_EMAIL` | Email admin default | `admin@example.com` |
| `ADMIN_PASSWORD` | Password admin default | `admin123` |

## Endpoint API

### **Authentication**
| Method | Endpoint | Deskripsi |
|---------|-----------|-----------|
| `POST` | `/register` | Register user baru |
| `POST` | `/login` | Login dan dapatkan JWT token |

### **Lapangan (Fields)**
| Method | Endpoint | Akses | Deskripsi |
|---------|-----------|--------|-----------|
| `GET` | `/fields` | Public | Daftar semua lapangan |
| `GET` | `/fields/:id` | Public | Detail lapangan |
| `POST` | `/fields` | Admin | Tambah lapangan |
| `PUT` | `/fields/:id` | Admin | Update lapangan |
| `DELETE` | `/fields/:id` | Admin | Hapus lapangan |

### **Booking**
| Method | Endpoint | Akses | Deskripsi |
|---------|-----------|--------|-----------|
| `POST` | `/bookings` | User | Buat booking baru |
| `GET` | `/bookings` | User/Admin | Lihat daftar booking |
|        | *(otomatis)* | | Validasi waktu agar tidak overlap |

### **Payment (Mock)**
| Method | Endpoint | Akses | Deskripsi |
|---------|-----------|--------|-----------|
| `POST` | `/payments` | User | Ubah status booking jadi `paid` |

## Contoh JSON

### Register
```json
{
  "name": "User 1",
  "email": "user1@example.com",
  "password": "123456"
}
```

### Login (Response)
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."
}
```

### Create Booking
```json
{
  "field_id": 1,
  "start_time": "2025-10-25T10:00:00Z",
  "end_time": "2025-10-25T12:00:00Z"
}
```

## Panduan Postman

1. Jalankan request **`/login`**
2. Copy token JWT dari response
3. Tambahkan ke header:
   ```
   Authorization: Bearer <token>
   ```
4. Import file `BookingFieldsApp.postman_collection.json` dari repo ini untuk testing

## Docker Commands Singkat

| Tujuan | Perintah |
|---------|-----------|
| Build & Run container | `docker compose up --build` |
| Jalankan di background | `docker compose up -d` |
| Stop container | `docker compose down` |
| Lihat log app | `docker logs -f booking-fields-app-app` |

## Lisensi

MIT License © 2025 — Created by Khoirul Azkiya  
Project ini dibuat untuk kebutuhan take-home test dan pengembangan backend API berbasis Golang.