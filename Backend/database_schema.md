# Database Schema & Business Logic - Internity

## 1. Deskripsi Sistem
Internity adalah sistem manajemen magang berbasis web.
Terdapat 3 role pengguna:
- **Admin**: Mengelola pengguna dan membuat berita/informasi (CRUD).
- **Mentor**: Melihat laporan kehadiran dan mengunduh rekap (Excel).
- **Intern**: Melakukan absensi via scan QR, mengisi form jobdesk, dan mengajukan izin/sakit.

---

## 2. Struktur Tabel Database

### Tabel `users`
Menyimpan data pengguna dan hak akses.
- `id` (Primary Key, Auto Increment/UUID)
- `name` (Varchar 255, Not Null)
- `email` (Varchar 255, Unique, Not Null)
- `role` (Enum: 'admin', 'mentor', 'intern', Not Null)
- `avatar` (Varchar 255, Nullable, path foto profil)
- `internship_start_date` (Date, Nullable)
- `internship_end_date` (Date, Nullable)
- `is_active` (Boolean, Default: true)
- `created_at` (Timestamp)
- `updated_at` (Timestamp)

### Tabel `attendances`
Menyimpan data absensi harian dan laporan jobdesk.
- `id` (Primary Key, Auto Increment/UUID)
- `user_id` (Foreign Key -> users.id, On Delete Cascade)
- `date` (Date, Not Null)
- `status` (Enum: 'hadir', 'izin', 'sakit', 'alfa', Not Null)
- `clock_in` (Time, Nullable)
- `clock_out` (Time, Nullable)
- `latitude_in` (Decimal 10,8, Nullable)
- `longitude_in` (Decimal 11,8, Nullable)
- `latitude_out` (Decimal 10,8, Nullable)
- `longitude_out` (Decimal 11,8, Nullable)
- `job_description` (Text, Nullable)
- `documentation_file` (Varchar 255, Nullable, path file bukti)
- `leave_reason` (Text, Nullable)
- `medical_letter_file` (Varchar 255, Nullable, surat dokter jika sakit)
- `created_at` (Timestamp)
- `updated_at` (Timestamp)

### Tabel `news`
Menyimpan berita/informasi dari perusahaan.
- `id` (Primary Key, Auto Increment/UUID)
- `title` (Varchar 255, Not Null)
- `content` (Text, Not Null)
- `image_path` (Varchar 255, Nullable)
- `author_id` (Foreign Key -> users.id, On Delete Cascade)
- `created_at` (Timestamp)
- `updated_at` (Timestamp)

### Tabel `email_otps`
Menyimpan kode OTP sementara untuk otentikasi login.
- `id` (Primary Key, Auto Increment/UUID)
- `email` (Varchar 255, Not Null)
- `otp_code` (Varchar 6, Not Null)
- `expires_at` (Timestamp, Not Null)
- `is_used` (Boolean, Default: false)
- `created_at` (Timestamp)

---

## 3. Aturan Logika Bisnis (Untuk referensi pembuatan Controller)

### A. Otentikasi (Login OTP)
1. User input email. Cek apakah email ada di tabel `users` dan `is_active = true`.
2. Jika ada, generate 6 digit angka random, simpan ke `email_otps` dengan masa aktif 5 menit.
3. Saat verifikasi, cek kesesuaian email dan OTP, serta pastikan `expires_at` belum lewat dan `is_used = false`.
4. Jika valid, update `is_used = true` dan berikan token otentikasi (JWT).

### B. Absensi (Clock In / Out)
1. **Clock In**: Menyimpan waktu saat ini ke `clock_in` beserta koordinat latitude & longitude.
2. **Clock Out**: Validasi waktu server harus >= pukul 17:00. Jika belum jam 17:00, tolak request.

### C. Batas Waktu Jobdesk (Cut-off jam 00:00)
1. Intern dapat mengupdate `job_description` dan `documentation_file`.
2. Backend WAJIB mengecek `date` dari absensi tersebut. Jika waktu server saat request masuk sudah berbeda hari (sudah melewati jam 23:59 pada tanggal tersebut), tolak request dengan error status 400/403.

### D. Auto "Alfa"
1. Terdapat task/cron job yang berjalan setiap hari pukul 23:59.
2. Cari semua user dengan `role = 'intern'` dan `is_active = true`.
3. Jika intern tersebut tidak memiliki record di tabel `attendances` pada hari itu, otomatis buat record baru dengan `status = 'alfa'`.