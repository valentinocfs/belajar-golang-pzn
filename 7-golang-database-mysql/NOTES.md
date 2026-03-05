### Package Database
Package database digunakan untuk mengelola koneksi ke database, dalam contoh ini kita akan menggunakan MySQL sebagai database yang akan kita gunakan.

### Database Driver
Untuk bisa menggunakan database MySQL di Golang, kita perlu menggunakan database driver yang sesuai dengan database yang kita gunakan. Untuk MySQL, kita bisa menggunakan driver bernama "github.com/go-sql-driver/mysql".

### Membuat Koneksi ke Database
Untuk membuat koneksi ke database, kita bisa menggunakan fungsi sql.Open() yang disediakan oleh package database/sql. Fungsi ini membutuhkan dua parameter, yaitu driver name dan data source name (DSN). Driver name adalah nama dari database driver yang kita gunakan, sedangkan data source name adalah string yang berisi informasi tentang koneksi ke database, seperti username, password, host, port dan nama database. Contoh DSN untuk MySQL adalah "username:password@tcp(host:port)/dbname".

### Database Pooling
Database pooling adalah teknik untuk mengelola koneksi ke database dengan cara membuat pool koneksi yang bisa digunakan secara bersamaan oleh beberapa goroutine. Dengan menggunakan database pooling, kita bisa meningkatkan performa aplikasi karena kita tidak perlu membuat koneksi baru setiap kali ada request yang masuk, melainkan kita bisa menggunakan koneksi yang sudah ada di pool.

- SetMaxIdleConns()     = jumlah koneksi minimal yang dibuat
- SetMaxOpenConns()     = jumlah koneksi maksimal yang dibuat
- SetConnMaxLifetime()  = berapa lama koneksi yang sudah tidak digunakan akan dihapus
- SetConnMaxIdleTime()  = berapa lama koneksi boleh digunakan sebelum dihapus

### Querying Data
- Exec()   = digunakan untuk menjalankan query yang tidak mengembalikan hasil, seperti INSERT, UPDATE, DELETE
- Query()  = digunakan untuk menjalankan query yang mengembalikan hasil, seperti SELECT

Exec.Context() dan QueryContext() adalah versi context-aware dari Exec() dan Query(), yang memungkinkan kita untuk menggunakan context dalam menjalankan query, sehingga kita bisa membatalkan query jika diperlukan.

### Tipe Data Column
Ketika kita melakukan query ke database, hasil yang kita dapatkan akan berupa tipe data yang sesuai dengan tipe data yang ada di database. Misalnya, jika kita memiliki kolom dengan tipe data INT, maka hasil yang kita dapatkan akan berupa tipe data int di Golang. Namun, jika kita memiliki kolom dengan tipe data NULLABLE, maka hasil yang kita dapatkan akan berupa tipe data sql.NullInt64, yang merupakan struct yang memiliki field Int64 dan Valid.

Tipe Data Database        | Tipe Data Golang
--------------------------|------------------
VARCHAR, CHAR, TEXT       | string
INT, BIGINT               | int, int64
FLOAT, DOUBLE             | float32, float64
BOOLEAN                   | bool
DATE, DATETIME, TIMESTAMP | time.Time

### SQL Injection
SQL Injection adalah sebuah teknik serangan yang memanfaatkan celah keamanan pada aplikasi yang menggunakan database, dengan cara menyisipkan kode SQL ke dalam input yang diberikan oleh user. Untuk mencegah SQL Injection, kita bisa menggunakan prepared statement atau parameterized query, yang memungkinkan kita untuk memisahkan antara kode SQL dan data yang diberikan oleh user, sehingga kode SQL tidak akan dieksekusi jika data yang diberikan oleh user mengandung kode SQL. 

### SQL With Parameter
SQL With Parameter adalah teknik untuk menjalankan query dengan menggunakan parameter, yang memungkinkan kita untuk memisahkan antara kode SQL dan data yang diberikan oleh user, sehingga kode SQL tidak akan dieksekusi jika data yang diberikan oleh user mengandung kode SQL. Contoh penggunaan SQL With Parameter adalah sebagai berikut:

```go 
query := "SELECT username FROM users WHERE username = ?"
rows, err := db.Query(query, userID)
```

### Auto Increment
Untuk mendapatkan last insert id setelah melakukan query INSERT, kita bisa menggunakan fungsi LastInsertId() yang disediakan oleh package database/sql. Fungsi ini akan mengembalikan id terakhir yang diinsert ke database, yang biasanya digunakan untuk mendapatkan id dari record yang baru saja diinsert ke database, terutama jika kita menggunakan auto increment pada kolom id. Contoh penggunaan LastInsertId() adalah sebagai berikut:

```go
result, err := db.Exec("INSERT INTO users (username) VALUES (?)", username)
if err != nil {
    panic(err)
} 
lastInsertID, err := result.LastInsertId()
if err != nil {
    panic(err)
} 
fmt.Println("Last Insert ID:", lastInsertID)
```

### Prepared Statement
Prepared statement adalah teknik untuk menjalankan query dengan menggunakan parameter, yang memungkinkan kita untuk memisahkan antara kode SQL dan data yang diberikan oleh user, sehingga kode SQL tidak akan dieksekusi jika data yang diberikan oleh user mengandung kode SQL. Prepared statement juga memiliki kelebihan dalam hal performa, karena query yang sudah diprepare akan disimpan di database dan bisa digunakan kembali tanpa perlu melakukan parsing ulang setiap kali dijalankan.

### Transaction
Database Transaction adalah sebuah mekanisme untuk menjalankan beberapa query secara atomik, artinya jika salah satu query gagal maka semua query yang dijalankan dalam transaction tersebut akan dibatalkan. Transaction biasanya digunakan untuk memastikan konsistensi data, terutama ketika kita melakukan operasi yang melibatkan beberapa tabel atau beberapa record sekaligus.

### Repository Pattern
Repository pattern adalah sebuah pola desain yang digunakan untuk memisahkan antara logika bisnis dengan logika akses data, sehingga kita bisa mengubah implementasi akses data tanpa perlu mengubah logika bisnis. Repository pattern biasanya digunakan untuk membuat aplikasi yang lebih modular dan mudah untuk diuji, karena kita bisa membuat mock repository untuk melakukan unit testing pada logika bisnis tanpa perlu terhubung ke database secara langsung.

#### Contoh Struktur Repository Pattern

- Repository
  - UserRepository
    - GetByID(id int) (*User, error)
    - Create(user *User) error
    - Update(user *User) error
- Entity
  - User (struct)
    - ID int
    - Username string
- Service
  - UserService
    - GetUserByID(id int) (*User, error)
    - CreateUser(user *User) error
    - UpdateUser(user *User) error
- Handler
  - UserHandler
    - GetUserByID(w http.ResponseWriter, r *http.Request)
    - CreateUser(w http.ResponseWriter, r *http.Request)
    - UpdateUser(w http.ResponseWriter, r *http.Request)
- Main
  - main.go
    - db, err := sql.Open("mysql", dsn)
    - userRepository := NewUserRepository(db)
    - userService := NewUserService(userRepository)
    - userHandler := NewUserHandler(userService)
    - http.HandleFunc("/users/{id}", userHandler.GetUserByID)
    - http.HandleFunc("/users", userHandler.CreateUser)
    - http.HandleFunc("/users/{id}", userHandler.UpdateUser)