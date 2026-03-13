## Database Migration
Database Migration merupakan mekanisme untuk melakukan tracking perubahan pada database schema dari awal dibuat sampai perubahan terakhir yang dilakukan.

### Golang Migrate
Golang Migrate adalah sebuah library yang digunakan untuk melakukan database migration pada aplikasi Golang. Golang Migrate mendukung banyak driver database seperti PostgreSQL, MySQL, SQLite, MongoDB, Cassandra dan lainnya.

### Installation
> go install -tags "mysql,postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

### Membuat Database Migration
> migrate create -ext sql -dir db/migrations "add_users_table"

### Migration Up and Down
> migrate up = menjalankan migrasi yang ingin kita ubah ke database
> migrate down = menjalankan migrasi yang ingin dihapus dari database

### Menjalankan Migration
> migrate -database "mysql://user:password@localhost:5432/dbname" -path folder up

### Migration State
Golang Migrate menyimpan state migrasi pada database, sehingga migrasi dapat dijalankan secara terurut dan tidak ada migrasi yang dijalankan lebih dari sekali. Semua informasi migrasi disimpan pada tabel `schema_migrations`.

### Rollback Migration
Rollback migration adalah proses mengembalikan migrasi ke versi sebelumnya. Golang Migrate menyimpan state migrasi pada database, sehingga rollback migration dapat dilakukan dengan mudah.
> migrate -database "mysql://user:password@localhost:5432/dbname" -path folder down

### Migrasi Ke Versi Tertentu
Untuk kasus migrasi ke versi tertentu, kita dapat menggunakan perintah `up` dengan argumen versi yaitu jumlah migrasi yang ingin dijalankan.
> migrate -database "mysql://user:password@localhost:5432/dbname" -path folder (up/down) 2

### Dirty State
Jika terlanjur migrasi terjadi secara tidak terkendali dan database berada dalam keadaan "dirty", kita dapat menghapus tabel `schema_migrations` pada database dan menjalankan migrasi ulang. Pada kasus ini, kita harus memperbaiki secara manual migrasi yang telah dijalankan sebelumnya.

- Memperbaiki Migrasi Yang Gagal
- Mengubah Versi Yang Gagal
> migrate -database "mysql://user:password@localhost:5432/dbname" -path folder force version
- Migrate ulang
> migrate -database "mysql://user:password@localhost:5432/dbname" -path folder up

### Migrasi File Tertentu
Untuk menjalankan migrasi file tertentu, kita dapat menggunakan argumen `-file` dengan path ke file migrasi yang ingin dijalankan.
> migrate -database "mysql://user:password@localhost:5432/dbname" -path folder up -file path/to/migration/file.sql
