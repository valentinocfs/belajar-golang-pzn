### Viper
Viper adalah library yang digunakan untuk manajemen konfigurasi aplikasi.

### Cara Membaca Viper
```go

config := viper.New()
config.SetConfigName("config")
config.SetConfigType("json")
config.AddConfigPath(".")

config.ReadInConfig()

name := config.GetString("name")
address := config.GetString("address")
port := config.GetInt("port")

println(name)
println(address)
println(port)

```

### Cara Set Viper
```go
config := viper.New()
config.Set("name", "Golang Viper")
config.Set("address", "localhost")
config.Set("port", 8080)

name := config.GetString("name")
address := config.GetString("address")
port := config.GetInt("port")

println(name)
println(address)
println(port)

```

### Environment Variable
Kadang saat menjalankan aplikasi, kita tidak ingin menggunakan file konfigurasi, melainkan menggunakan environment variable. Secara default, viper tidak bisa membaca environment variable. Kita perlu mengaktifkan fitur ini dengan menggunakan method `AutomaticEnv()`.

### Fitur Lainnya
Jenis file konfigurasi yang bisa dibaca oleh viper:
- HCL
- TOML
- YAML
- JSON
- ENV
- Consul
- Etcd
- Redis
- Zookeeper