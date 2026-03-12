### Golang Generics
Golang generics adalah fitur baru yang memungkinkan penggunaan tipe data generik.

### Type Parameters
Type parameters adalah tipe data generik yang digunakan pada saat runtime. Type parameters menggunakan tanda [] untuk mendefinisikan tipe data generik.

### Type Constraints
Type constraints adalah batasan tipe data yang digunakan pada saat runtime. Type constraints menggunakan tanda interface{} untuk mendefinisikan batasan tipe data.

### Multiple Type Parameters
Multiple type parameters adalah fitur yang memungkinkan penggunaan lebih dari satu type parameter pada saat runtime. Multiple type parameters menggunakan tanda , (koma) untuk mendefinisikan type parameter yang lebih dari satu. Nama type parameter harus unik.

### Comparable Types
Comparable types adalah tipe data yang dapat dibandingkan menggunakan operator == dan !=. Comparable types menggunakan tanda interface{} untuk mendefinisikan batasan tipe data.

### Type Parameter Inheritance
Type parameter inheritance adalah fitur yang memungkinkan type parameter untuk mewarisi batasan tipe data dari type parameter lainnya. Type parameter inheritance menggunakan tanda interface{} untuk mendefinisikan batasan tipe data.

### Type Sets
Type sets adalah kumpulan tipe data yang digunakan pada saat runtime. Type sets menggunakan tanda interface{} untuk mendefinisikan batasan tipe data.
```go
type MyType[T interface{}] struct {
	int | int8 | float64
}
```

### Type Declaration
Type declaration adalah fitur yang memungkinkan penggunaan type alias pada saat runtime. Type declaration menggunakan tanda type untuk mendefinisikan type alias.

### Type Appromaximation

### Type Inference

### Generic Type

### Generic Struct

### Generic Interface

### In Line Type Constraints
```
interface{ int | int8 | float64 }
```
