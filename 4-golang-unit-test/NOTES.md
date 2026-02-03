# GOLANG UNIT TEST

## CARA MENJALANKAN UNIT TEST

- `go test` => Mengcompile dan menjalankan unit test
- `go test -v` => Mengcompile dan menjalankan unit test dengan verbose
- `go test -run {nama test function}` => Mengcompile dan menjalankan unit test dengan run nama test function 

- `go test ./...` => Mengcompile dan menjalankan unit test di semua package

## CARA MENGGAGALKAN UNIT TEST

- `t.Fail()`    => Mencatat error tapi tidak menghentikan eksekusi unit test
- `t.FailNow()` => Mencatat error dan menghentikan eksekusi unit test
- `t.Error("Pesan error")` => Mencatat error dan tidak menghentikan eksekusi unit test
- `t.Fatalf("Pesan error")` => Mencatat error dan menghentikan eksekusi unit test
- `t.Errorf("Pesan error")` => Mencatat error dan tidak menghentikan eksekusi unit test

> Lebih disarankan menggunakan t.Error dan t.Fatal

## ASSERTION

- Testify => Library assertion untuk golang
- https://github.com/stretchr/testify/assert
- https://github.com/stretchr/testify/require

## SKIP TEST

- `t.Skip("Pesan skip")` => Melewati eksekusi unit test

## BEFORE & AFTER UNIT TEST

- `m.Run()` => Menjalankan function sebelum dan sesudah test selesai

## SUB TEST

- `t.Run("Nama test", func(t *testing.T) { ... })` => Membuat sub test
- Menjalankan sub test:
  - `go test -v -run=TestSubTest/Test_1`
  - `go test -v -run /Test_1`

## TABLE TEST

- `[]struct { ... }` => Membuat array of struct
- `t.Run("Nama test", func(t *testing.T) { ... })` => Membuat sub test


## MOCK

- Mock => Mengganti function yang di test dengan function yang sudah di test
- Mocking => Mengganti function yang di test dengan function yang sudah di test
- https://github.com/stretchr/testify


# BENCHMARK

- `go test -v -run=TestNotMatch -bench=.` => Menjalankan semua benchmark
- `go test -v -run=TestNotMatch -bench={namabenchmark}` => Menjalankan benchmark dengan nama benchmark
- `go test -v -run=TestNotMatch -bench=. -benchmem` => Menjalankan semua benchmark dengan memori

# SUB BENCHMARK

- `t.Run("Nama benchmark", func(b *testing.B) { ... })` => Membuat sub benchmark

# BENCHMARK TABLE

- `[]struct { ... }` => Membuat array of struct
- `t.Run("Nama benchmark", func(b *testing.B) { ... })` => Membuat sub benchmark