## BookEvent API

**Selamat datang di BookEvent API!**

API ini memungkinkan Anda untuk:

- **Mengelola pengguna:**
    - Membuat akun pengguna baru.
    - Mengambil informasi pengguna berdasarkan ID.
    - Mengubah data pengguna.
    - Menghapus pengguna.
- **(Fitur tambahan)**

**Dokumentasi:**

- **Endpoint:**
    - `/users/signup`: Buat akun pengguna baru.
    - `/users/login`: Autentikasi pengguna dan dapatkan token.
    - `/users`: Dapatkan semua pengguna.
    - `/users/:id`: Dapatkan pengguna berdasarkan ID.
    - `/users/:id/update`: Ubah data pengguna.
    - `/users/:id/delete`: Hapus pengguna.

**Contoh Penggunaan:**

- **Membuat akun pengguna:**

```
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newemail@example.com",
    "password": "password123"
  }' \
  "https://api.example.com/users/signup"
```

- **Mengambil informasi pengguna:**

```
curl -X GET \
  -H "Authorization: Bearer <token>" \
  "https://api.example.com/users/123"
```

**Persyaratan:**

- Go v1.18+
- PostgreSQL database

**Instalasi:**

1. Clone repository ini.
2. Jalankan `go mod download`.
3. Jalankan `go run main.go`.

**Catatan:**

- Ini adalah contoh project sederhana. Anda dapat menambahkan fitur dan fungsionalitas sesuai kebutuhan.
- Pastikan Anda telah mengkonfigurasi database PostgreSQL dengan benar.
- Untuk keamanan, gunakan HTTPS dan token autentikasi saat menggunakan API.

**Fitur Tambahan:**

- **Manajemen Event:**
    - Buat, edit, dan hapus event.
    - Reservasi dan daftar tunggu untuk event.
- **Manajemen Kategori:**
    - Buat, edit, dan hapus kategori event.

