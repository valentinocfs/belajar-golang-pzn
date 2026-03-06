### ServeMux
ServeMux adalah sebuah struct yang digunakan untuk mengatur routing pada aplikasi web di Go. ServeMux memungkinkan kita untuk menentukan handler yang akan menangani permintaan HTTP berdasarkan path yang diberikan. ServeMux juga menyediakan fitur untuk menangani path yang memiliki parameter, seperti /users/{id}, sehingga kita bisa menangani permintaan HTTP dengan path yang dinamis.

### HTTP Test
HTTP Test adalah sebuah package yang digunakan untuk melakukan testing pada aplikasi web di Go. HTTP Test menyediakan fitur untuk membuat request HTTP dan memeriksa response yang diberikan oleh handler, sehingga kita bisa memastikan bahwa handler yang kita buat sudah bekerja dengan benar.

### Query Parameter
Query parameter adalah sebuah parameter yang ditambahkan pada URL setelah tanda tanya (?), yang digunakan untuk mengirimkan data tambahan pada permintaan HTTP. Query parameter biasanya digunakan untuk mengirimkan data yang tidak terlalu penting atau data yang bersifat opsional, seperti filter, sorting, atau pagination.

### Header
writer.Header.Add("Content-Type", "application/json")
writer.Header.Set("Content-Type", "application/json")

### Form Post
Request.ParseForm()
Request.PostForm.Get("name")

### Cookie
Cookie adalah sebuah mekanisme yang digunakan untuk menyimpan data pada sisi klien, yang dapat digunakan untuk mengidentifikasi pengguna atau menyimpan preferensi pengguna. Cookie biasanya digunakan untuk menyimpan data yang bersifat sementara, seperti session ID atau token autentikasi.

### FileServer
FileServer adalah sebuah handler yang digunakan untuk menyajikan file statis pada aplikasi web di Go. FileServer memungkinkan kita untuk menyajikan file dari direktori tertentu, sehingga kita bisa menyajikan file seperti gambar, CSS, atau JavaScript pada aplikasi web kita. FileServer juga menyediakan fitur untuk menangani path yang memiliki parameter, seperti /static/{filename}, sehingga kita bisa menyajikan file dengan nama yang dinamis.

### ServeFile
ServeFile adalah fungsi yang digunakan untuk menyajikan file statis melalui HTTP yang merupakan bagian dari paket http dan memungkinkan untuk mengirimkan file kepada klien sebagai respon terhadap permintaan HTTP. 

### Template
{{..}}

### Template Data, Action, Layout, Function, & Caching
#### Data
{{.Title}} {{.Name}} {{.Address.Street}}
#### Action
- Pengkondisian
{{if.Value}} T1 {{end}}
{{if.Value}} T1 {{else}} T2 {{end}}
{{if.Value1}} T1 {{else if.Value2}} T2 {{else}} T3 {{end}}
- Perbandingan
{{if ge .FinalValue 80}} {{end}}
- Range
{{ range $index, $element := .Value}} {{else}} {{end}}
- With (nested struct)
{{with.Value}} T1 {{end}}
- Komentar
{{/* This is comment in go template */}}
#### Layout
- Import Template
{{ template "nama" }}
{{ template "nama". }}
{{ template "nama".Value }}
#### Function
- Method / Function on Struct
{{.FunctionName}}
{{.FunctionName "param1", "param2"}} 
- Global Function
Semua function built in bisa dipanggil atau Global function buatan sendiri
- Function Pipelines
{{ sayHello .Name | upper }}
#### Caching
Menyimpan data template caching sehingga kita tidak perlu melakukan parsing lagi

### XSS Cross Site Scripting
- Secara default golang otomatis menambahkan ketika render: 
{{ urlescaper | attrescaper}} {{ htmlescaper }}
- Jika ingin merender elemen html/css/js bisa menggunakan
> template.HTML()
> template.CSS()
> template.JS()

### Redirect
http.Redirect()

### Upload & Download File
- MultiPart
  - multipart.File
  - multipart.FileHeader
- ServeFile = Untuk Preview File
- Content Disposition = Untuk Download File

### Middleware
Tidak ada istilah middleware di golang web namun bentuknya dalam bentuk interface
- Bisa digunakan sebagai error handler

### Routing Library
ServeMux mempunyai fitur yang terbatas, mostly developer golang menggunakan library untuk routing seperti:
- HttpRouter
- GorilaMux