### HTTP Router
Library open source untuk HTTP handler di Golang
Terkenal karena kecepatannya dan sifat yang minimalis

### Method
GET, POST, PUT, DELETE 

### Params
Params digunakan untuk menyimpan parameter yang dikirim dari client di parameter URL

### Route Pattern
- Named Paramater
Pattern               /user/:user
/user/john            match
/user/john/profile    no match
/user/                no match
- Catch All Parameter
Pattern               /src/*filepath
/src/                 no match
/src/somefile         match
/src/subdir/somefile  match

### Serve File
ServeFile adalah metode yang digunakan untuk melayani file statis, seperti gambar, CSS, dan JavaScript, kepada klien. Metode ini mengambil path file sebagai parameter dan secara otomatis mengatur header sehingga file dapat diunduh oleh pengguna dengan cara yang benar.

### Panic Handler
PanicHandler: func(http.ResponseWriter, *http.Request, interface{})

### Not Found Handler
NotFoundHandler : http.Handler

### Method Not Allowed Handler
MethodNotAllowed : http.Handler

### Middleware
HttpRouter hanyalah library untuk http router saja, tidak ada fitur lain selain router, dan karena router merupakan implementasi dari http.Handler, jadi untuk middleware kita bisa membuat dan custom sendiri 