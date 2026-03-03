### Parallel Programming
Parallel Programming adalah kemampuan untuk menjalankan beberapa program secara bersamaan.

### Process vs Thread
Process adalah program yang sedang berjalan.
Thread adalah bagian dari proses yang sedang berjalan.

### Parallelism vs Concurrency
Parallelism adalah kemampuan untuk menjalankan beberapa program secara bersamaan.
Concurrency adalah kemampuan untuk menjalankan beberapa program secara bergantian.

### Goroutine
Goroutine adalah thread yang ringan yang dibuat oleh Go yang dikelola oleh Go Runtime.
Ukuran default goroutine adalah 2KB.
Namun tidak seperti thread yang berjalan parallel, goroutine berjalan secara concurrency.

### Channel
Channel adalah jalur komunikasi antar goroutine.
Channel mirip seperti konsep async/await di javascript.
Channel harus di close jika tidak digunakan, atau bisa menyebabkan memory leak

### Buffer Channel
Buffer Capacity 
cap = melihat panjang buffer (capacity)
len = melihat jumlah data di buffer

### Range Channel
Range channel digunakan untuk membaca data dari channel sampai channel tersebut di close.

### Select Channel
Select channel digunakan untuk membaca data dari beberapa channel sekaligus.

### Default Select Channel
Default select channel digunakan untuk membaca data dari beberapa channel sekaligus, jika tidak ada data yang masuk maka akan menjalankan default case.

### Race Condition
Race condition adalah kondisi dimana beberapa goroutine mencoba untuk mengakses data yang sama secara bersamaan, sehingga bisa menyebabkan data menjadi tidak konsisten.

### Mutex
Mutex (Mutual Exclusion) adalah mekanisme untuk mengunci data agar hanya satu goroutine yang bisa mengakses data tersebut pada saat yang sama.

### RWMutex
RWMutex (Read-Write Mutex) adalah mekanisme untuk mengunci data agar hanya satu goroutine yang bisa mengakses data tersebut pada saat yang sama, namun memungkinkan beberapa goroutine untuk membaca data tersebut secara bersamaan.

### Deadlock
Deadlock adalah kondisi dimana beberapa goroutine saling menunggu untuk mengakses data yang sama, sehingga tidak ada goroutine yang bisa melanjutkan eksekusi.

### WaitGroup
WaitGroup adalah mekanisme untuk menunggu beberapa goroutine selesai sebelum melanjutkan eksekusi.

### Once
Once adalah mekanisme untuk menjalankan sebuah fungsi hanya sekali, walaupun dipanggil berkali-kali.

### Pool
Pool adalah mekanisme untuk mengelola goroutine yang sudah selesai, sehingga bisa digunakan kembali untuk menjalankan goroutine yang baru.

### Cond
Cond adalah mekanisme untuk menunggu sebuah kondisi terpenuhi sebelum melanjutkan eksekusi. Cond membutuhkan mutex untuk mengunci data yang akan di cek kondisinya, dan juga membutuhkan channel untuk menunggu kondisi terpenuhi.

### Atomic
Atomic adalah mekanisme untuk mengakses data secara atomik, sehingga bisa menghindari race condition tanpa harus menggunakan mutex. Atomic hanya bisa digunakan untuk tipe data yang sudah disediakan oleh package sync/atomic, seperti int32, int64, uint32, uint64, uintptr, dan unsafe.Pointer.

### Timer
Timer adalah mekanisme untuk menjalankan sebuah fungsi setelah beberapa waktu tertentu. Timer bisa digunakan untuk membuat timeout, atau untuk menjalankan sebuah fungsi secara periodik.

### Ticker
Ticker adalah mekanisme untuk menjalankan sebuah fungsi secara periodik. Ticker akan menjalankan fungsi setiap interval waktu tertentu, sampai ticker tersebut dihentikan.

### GOMAXPROCS
GOMAXPROCS adalah variabel lingkungan yang digunakan untuk mengatur jumlah CPU yang digunakan oleh Go Runtime untuk menjalankan goroutine. Secara default, GOMAXPROCS akan diatur sesuai dengan jumlah CPU yang tersedia di sistem. Namun, kita bisa mengatur GOMAXPROCS secara manual untuk meningkatkan performa aplikasi kita, terutama jika aplikasi kita memiliki banyak goroutine yang berjalan secara parallel.