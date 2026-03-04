### Context
Context merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout dan sinyal deadline
Biasanya dibuat per request (misal setiap ada request masuk ke server web melalui http request)
Context digunakan untuk mempermudah kita meneruskan value, dan sinyal antar proses

### Membuat Context
- context.Background()
- context.TODO()

### Parent dan Child Context
Context menganut konsep parent dan child, artinya kita bisa membuat child context dari context yang sudah ada
Context bersifat immutable, artinya ketika kita membuat child context dari parent context, maka child context tidak akan merubah parent context

### Context with Value
Context with value digunakan untuk menyimpan value di dalam context, value ini bisa diambil di child context

### Context with Cancel
Context with cancel digunakan untuk membuat context yang bisa dibatalkan, ketika context dibatalkan maka semua child context yang dibuat dari context tersebut juga akan dibatalkan 

### Context with Timeout
Context with timeout digunakan untuk membuat context yang memiliki batas waktu, ketika waktu habis maka context akan dibatalkan secara otomatis

### Context with Deadline
Context with deadline digunakan untuk membuat context yang memiliki batas waktu tertentu, berbeda dengan context with timeout yang memiliki batas waktu relatif, context with deadline memiliki batas waktu absolut (misal tanggal 1 Januari 2026 jam 12:00:00)