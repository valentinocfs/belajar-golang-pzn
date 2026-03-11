### Dependency Injection
Dependency Injection merupakan sebuah teknik dimana sebuah object menerima object lain yang dibutuhkan (dependencies) ketika pembuatan object itu sendiri.
Biasanya object yang menerima dependencies disebut client, proses mengirim dependencies disebut injection, dan object yang dikirim disebut provider.

- Function Sebagai Constructor
```go
func NewUserService(userStore user.Store) *UserService {
	return &UserService{
		userStore: userStore,
	}
}
```

Namun jika dilakukan secara manual dengan skala aplikasi yang besar, maka akan ada potensi terjadinya siklus dependensi (circular dependency). Maka dari itu, perlu adanya library yang membantu dalam implementasi Dependency Injection.

### Library Dependency Injection
- Google Wire
Google Wire adalah library yang membantu dalam implementasi Dependency Injection. Google Wire menggunakan reflection untuk melakukan dependency injection. Google Wire memiliki kelebihan yaitu tidak memerlukan runtime reflection. Google Wire memiliki kekurangan yaitu perlu adanya file wire yang harus dijalankan untuk menghasilkan file .pb.go.

### Provider (constructor)
Provider adalah function yang mengembalikan object yang akan diinject ke dalam client. 

### Injector
Injector adalah function yang akan di-generate oleh Google Wire. Injector akan mengembalikan object yang akan diinject ke dalam client. 

### Error pada Provider
Google Wire bisa mendeteksi jika terjadi error pada Provider. Jika terdapat error, secara otomatis akan mengembalikan error ketika kita melakukan dependency injection. Caranya sederhana, kita cukup buat di Provider return value kedua berupa error, dan di injectornya juga perlu ditambahkan return value kedua berupa error

### Injector Parameter
Injector Parameter merupakan parameter yang diberikan ke injector

### Multiple Binding
Multiple Binding merupakan kemampuan untuk mengikat beberapa variabel dalam satu pernyataan ataupun ekspresi.

### Provider Set
Provider Set digunakan untuk melakukan grouping provider

### Binding Interface
Binding interface yaitu memberitahu ke wire provider mana yang harus digunakan ketika interface tersebut dibutuhkan

### Struct Provider
Struct Provider merupakan struct yang bisa kita jadikan Provider

### Binding Values
Binding Values adalah kemampuan untuk mengikat nilai ke dalam Provider

### Struct Field Provider
Struct Field Provider merupakan kemampuan untuk mengikat field struct ke dalam Provider

### Cleanup Function
Cleanup Function adalah function yang akan dipanggil ketika object sudah tidak digunakan lagi