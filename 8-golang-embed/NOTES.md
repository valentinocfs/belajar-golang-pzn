### Embed Package
Package embed digunakan untuk menyisipkan file atau direktori ke dalam binary aplikasi Golang. Dengan menggunakan package embed, kita bisa menyisipkan file seperti HTML, CSS, JavaScript, gambar, atau file lainnya ke dalam aplikasi Golang, sehingga kita tidak perlu lagi mengandalkan file eksternal yang harus disertakan saat menjalankan aplikasi.

### Cara Menggunakan Embed Package
Untuk menggunakan package embed, kita perlu mengimport package embed terlebih dahulu, kemudian kita bisa menggunakan directive //go:embed untuk menyisipkan file atau direktori ke dalam aplikasi. Directive //go:embed harus diletakkan di atas deklarasi variabel yang akan digunakan untuk menyimpan file atau direktori yang disisipkan.
Contoh penggunaan embed package untuk menyisipkan file a.txt adalah sebagai berikut:

```go
package main
import (
  "embed"
  "fmt"
)

//go:embed files/a.txt
var a string
func main() {
  fmt.Println(a)
}
```

### Path Matcher
Path matcher adalah fitur yang disediakan oleh package embed untuk menyisipkan file atau direktori dengan menggunakan pola path tertentu. Dengan menggunakan path matcher, kita bisa menyisipkan file atau direktori yang memiliki nama atau ekstensi tertentu, sehingga kita tidak perlu menyisipkan file satu per satu. Contoh penggunaan path matcher untuk menyisipkan semua file dengan ekstensi .txt di dalam direktori files adalah sebagai berikut:

```go
package main
import (
  "embed"
  "fmt"
)

//go:embed files/*.txt
var files embed.FS

func main() {
  dir, _ := files.ReadDir("files")
  for _, d := range dir {
    fmt.Println(d.Name())
  }
}