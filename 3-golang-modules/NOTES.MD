# GO MODULES

## SEBELUM ADA GO MODULES

- Sebelum ada Go Modules, management dependency di Go itu sedikit merepotkan.
- Biasanya kita akan copy semua source code library yang kita butuhkan ke dalam folder project kita, hal ini membuat project kita menjadi bengkak karena penuh dengan library lain.
- Atau biasanya library orang lain kita taruh di dalam folder $GOPATH/src/nama_library, namun hal ini akan sulit jika ternyata beberapa aplikasi butuh versi library yang berbeda.

## GO MODULES

- Go Modules adalah management dependency di Go yang diperkenalkan pada Go versi 1.11.
- Go Modules memudahkan kita dalam management dependency di Go.
- Go modules itu seperti composer di PHP, yarn/npm di JavaScript, atau pip di Python.

## CARA MEMBUAT GO MODULES

- Untuk membuat go modules, kita perlu membuat file go.mod di dalam project kita.
- Caranya dengan menjalankan perintah `go mod init nama_module`.
- Contoh: `go mod init app-say-hello`

## MENGUBAH VERSI DEPENDENCY (MINOR & PATCH)

- Untuk mengubah versi dependency ubah go.mod manual, lalu jalankan perintah `go get`.
- Contoh: `go get`

## MENGUBAH VERSI DEPENDENCY (MAJOR)

- Untuk mengubah versi dependency ubah go.mod manual, komentar baris require yang lama dan tambahkan require yang baru dengan prefix `/v2` (sesuai versi yang ingin diubah), lalu jalankan perintah `go get nama_module/v2`.
- Contoh: `go get github.com/ProgrammerZamanNow/go-say-hello/v2`