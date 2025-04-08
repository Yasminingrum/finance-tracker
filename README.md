# Dokumentasi API Finance Tracker

## Pengantar

Dokumen ini berisi dokumentasi lengkap untuk API Finance Tracker. API ini memungkinkan pengguna melakukan pendaftaran akun, login, serta mengelola transaksi keuangan mereka.

## Autentikasi

API Finance Tracker menggunakan JWT (JSON Web Token) untuk autentikasi.

- **Metode**: Bearer Token
- **Header yang diperlukan**: 
  ```
  Authorization: Bearer <your_jwt_secret>
  ```
- Token dibuat menggunakan secret key dari file `.env` (JWT_SECRET)
- Token berisi ID pengguna yang digunakan untuk mengidentifikasi dan mengotorisasi pengguna

## Endpoints

### User (Pengguna)

#### 1. Registrasi Pengguna

Endpoint ini digunakan untuk mendaftarkan pengguna baru dalam sistem.

- **Path**: `/api/register`
- **Method**: POST
- **Auth**: Tidak perlu autentikasi
- **Headers**: Content-Type: application/json
- **Request Body**:
  ```json
  {
    "name": "Nama Lengkap",
    "email": "user@example.com",
    "password": "password_anda"
  }
  ```
  | Field    | Tipe   | Deskripsi                                          | Wajib |
  |----------|--------|---------------------------------------------------|-------|
  | name     | string | Nama lengkap pengguna, maksimal 100 karakter       | Ya    |
  | email    | string | Alamat email pengguna, harus unik, maks 100 karakter | Ya    |
  | password | string | Password pengguna, minimal 6 karakter              | Ya    |

- **Response**:
  - **201 Created**: Registrasi berhasil
    ```json
    {
      "message": "Registration successful"
    }
    ```
  - **400 Bad Request**: Email sudah terdaftar atau data tidak valid
    ```json
    {
      "error": "email already registered"
    }
    ```
    atau
    ```json
    {
      "error": "[pesan error validasi]"
    }
    ```
  - **500 Internal Server Error**: Error saat hashing atau penyimpanan ke database

#### 2. Login Pengguna

Endpoint ini digunakan untuk mengautentikasi pengguna dan mendapatkan token JWT.

- **Path**: `/api/login`
- **Method**: POST
- **Auth**: Tidak perlu autentikasi
- **Headers**: Content-Type: application/json
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "password_anda"
  }
  ```
  | Field    | Tipe   | Deskripsi                         | Wajib |
  |----------|--------|------------------------------------|-------|
  | email    | string | Alamat email yang sudah terdaftar  | Ya    |
  | password | string | Password akun                      | Ya    |

- **Response**:
  - **200 OK**: Login berhasil
    ```json
    {
      "message": "Login successful",
      "token": "jwt_token_anda"
    }
    ```
  - **401 Unauthorized**: Email atau password salah
    ```json
    {
      "error": "email not found"
    }
    ```
    atau
    ```json
    {
      "error": "incorrect password"
    }
    ```
  - **400 Bad Request**: Format data tidak valid
    ```json
    {
      "error": "[pesan error validasi]"
    }
    ```

### Transaction (Transaksi)

#### 1. Membuat Transaksi Baru

Endpoint ini digunakan untuk mencatat transaksi baru.

- **Path**: `/api/transactions`
- **Method**: POST
- **Auth**: Diperlukan autentikasi JWT
- **Headers**: 
  - Authorization: Bearer <jwt_token>
  - Content-Type: application/json
- **Request Body**:
  ```json
  {
    "type": "expense",
    "amount": 150000,
    "category": "Groceries",
    "note": "Belanja bulanan",
    "date": "2025-04-08"
  }
  ```
  | Field    | Tipe   | Deskripsi                                 | Wajib |
  |----------|--------|-------------------------------------------|-------|
  | type     | string | Jenis transaksi, hanya "income" atau "expense" | Ya    |
  | amount   | number | Jumlah uang dalam transaksi               | Ya    |
  | category | string | Kategori transaksi, maksimal 100 karakter | Tidak |
  | note     | string | Catatan tambahan tentang transaksi        | Tidak |
  | date     | string | Tanggal transaksi dalam format YYYY-MM-DD | Ya    |

- **Response**:
  - **201 Created**: Transaksi berhasil dibuat
    ```json
    {
      "message": "transaction successfully added"
    }
    ```
  - **400 Bad Request**: Format data tidak valid
    ```json
    {
      "error": "[pesan error validasi]"
    }
    ```
  - **401 Unauthorized**: Token tidak valid atau kadaluarsa
    ```json
    {
      "error": "unauthorized"
    }
    ```
  - **500 Internal Server Error**: Gagal menyimpan transaksi
    ```json
    {
      "error": "failed to create transaction"
    }
    ```

#### 2. Mendapatkan Semua Transaksi

Endpoint ini digunakan untuk mengambil semua transaksi milik pengguna yang sedang login.

- **Path**: `/api/transactions`
- **Method**: GET
- **Auth**: Diperlukan autentikasi JWT
- **Headers**: Authorization: Bearer <jwt_token>
- **Query Parameters**: Tidak ada

- **Response**:
  - **200 OK**: Daftar transaksi berhasil diambil
    ```json
    {
      "transactions": [
        {
          "id": 1,
          "user_id": 2,
          "type": "expense",
          "amount": 150000,
          "category": "Groceries",
          "note": "Belanja bulanan",
          "date": "2025-04-08T00:00:00Z",
          "created_at": "2025-04-08T14:12:00Z"
        },
        {
          "id": 2,
          "user_id": 2,
          "type": "income",
          "amount": 5000000,
          "category": "Salary",
          "note": "Gaji bulan April",
          "date": "2025-04-01T00:00:00Z",
          "created_at": "2025-04-01T10:00:00Z"
        }
      ]
    }
    ```
  - **401 Unauthorized**: Token tidak valid atau kadaluarsa
    ```json
    {
      "error": "unauthorized"
    }
    ```
  - **500 Internal Server Error**: Gagal mengambil data transaksi
    ```json
    {
      "error": "failed to retrieve data"
    }
    ```

## Status Code yang Digunakan

| Status Code | Deskripsi                                         |
|-------------|---------------------------------------------------|
| 200         | OK - Permintaan berhasil                          |
| 201         | Created - Resource berhasil dibuat                |
| 400         | Bad Request - Validasi gagal atau request invalid |
| 401         | Unauthorized - Autentikasi gagal                  |
| 404         | Not Found - Resource tidak ditemukan              |
| 500         | Internal Server Error - Kesalahan pada server     |

## Format Data

### Transaction

```json
{
  "id": 1,
  "user_id": 2,
  "type": "expense", // atau "income"
  "amount": 150000,
  "category": "Groceries",
  "note": "Belanja bulanan",
  "date": "2025-04-08T00:00:00Z",
  "created_at": "2025-04-08T14:12:00Z"
}
```

### User

```json
{
  "id": 1,
  "name": "Nama Lengkap",
  "email": "user@example.com",
  "created_at": "2025-04-01T10:00:00Z"
}
```

## Catatan Teknis

- Semua endpoint API harus diawali dengan `/api`
- API Finance Tracker menggunakan framework Gin untuk routing
- Autentikasi menggunakan JWT dengan payload berisi ID pengguna
- Password di-hash menggunakan bcrypt sebelum disimpan ke database
- Database menggunakan GORM sebagai ORM dengan model yang telah didefinisikan