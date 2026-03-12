### Logging
Log File adalah file yang berisikan informasi kejadian dari sebuah sistem. Logging adalah aksi menambah informasi ke log file. Loggin sudah menjadi standar industri untuk menampilkan informasi kejadian dari sebuah sistem. Logging bukan hanya untuk menampilkan informasi, tetapi juga untuk debugging dan troubleshooting.

### Ekosistem Logging
- Aplikasi (Mengirim Log Event)
- Log File
- Log Aggregator (Logstash)
- Log Database (Elasticsearch)
- Log Management

### Logging Library
Pada golang sebenarnya ada package log untuk melakukan logging namun fiturnya terbatas. Oleh karena itu, banyak library logging yang tersedia untuk golang seperti Logrus, Zap, Zerolog, dan lainnya.

### Logrus
#### Logger
Loger adalah struct utama pada Logrus untuk melakukan logging. Logger memiliki method-method untuk menambah log event seperti Info(), Error(), dan lainnya.

#### Level
Level adalah tingkat kepentingan dari log event. Level ini dapat digunakan untuk mengatur log yang akan ditampilkan.
- Trace = logger.Trace()
- Debug = logger.Debug()
- Info = logger.Info()
- Warn = logger.Warn()
- Error = logger.Error()
- Fatal = logger.Fatal()
- Panic = logger.Panic()

Trace dan Debug tidak dikeluar di console krn secara default levelnya lebih rendah dari Info. Untuk mengubah level, dapat menggunakan method SetLevel() pada logger.

#### Output
Output adalah tempat log event akan ditulis. Logrus mendukung output seperti console, file, dan lainnya. Output ini dapat digunakan untuk mengatur tempat log akan ditulis.

#### Formatter
Formatter adalah format log yang digunakan untuk menampilkan log event. Logrus mendukung format seperti JSON, plain text, dan lainnya. Formatter ini dapat digunakan untuk mengatur bagaimana log akan ditampilkan. Untuk mengubah formatter, dapat menggunakan method SetFormatter() pada logger.

#### Fields
Fields adalah data tambahan yang dapat ditambahkan ke log event. Fields ini dapat digunakan untuk memberikan konteks lebih lanjut tentang log event. Untuk menambah fields, dapat menggunakan method WithFields() pada logger.

#### Entry
Entry adalah sebuah struct yang digunakan untuk membuat log event. Entry ini dapat digunakan untuk menambah fields dan formatter ke log event. Untuk membuat entry baru, dapat menggunakan logrus.NewEntry()

#### Hook
Hook adalah struct yang digunakan untuk menambahkan hook ke logger. Hook ini dapat digunakan untuk melakukan sesuatu ketika log event terjadi, seperti mengirim log ke server atau mengirim email. Untuk menambah hook, dapat menggunakan method logger.AddHook() pada logger.

#### Singleton
Singleton adalah pattern yang digunakan untuk membuat hanya satu instance dari sebuah struct. Logrus menggunakan singleton untuk membuat logger yang dapat digunakan secara global.
