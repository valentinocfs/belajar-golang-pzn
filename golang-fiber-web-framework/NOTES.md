## Web Framework

### Web Framework Golang
- Fiber
- Gin
- Echo

### Fiber

### Configuration
fiber.Config{}

### Routing
fiber.App(path, handlers)

### Context
fiber.Ctx

### HTTP Request
ctx.Get(), ctx.Post(), ctx.Put(), ctx.Delete(), ctx.Cookies()

### Route Parameters
ctx.Params(), ctx.AllParams()

### Query Parameters
ctx.Query(), ctx.AllQuery()

### Request Form
ctx.FormValue(), ctx.AllFormValues()

### Multipart Form
ctx.MultipartForm(), ctx.FormFile()
ctx.SaveFile()

### Request Body
ctx.Body(), ctx.AllBody()

### Body Parser
ctx.BodyParser()  // v2
ctx.Bind().Body() // v3

### HTTP Response
ctx.Set()
ctx.Status()
ctx.Send()
ctx.XML()
ctx.JSON()
ctx.Redirect()
ctx.SendString()
ctx.Cookie()

### Download File
ctx.Download()

### Routing Group
fiber.Group(path, handlers)

### Static
- V2
```
app.Static(path, root)
```
- V3
```
static.Use("/public", static.New("", static.Config{
	FS:     os.DirFS("source"),
	Browse: true,
}))
app.Use("/static", static.New("./public/hello.html"))
app.Use("/", static.New("./public"))
```
- V3 embed
```
//go:embed path/to/files
var myfiles embed.FS
app.Get("/files*", static.New("", static.Config{
    FS:     myfiles,
    Browse: true,
}))
```
- V3 SPA
```
app.Use("/web", static.New("", static.Config{
    FS: os.DirFS("dist"),
}))
app.Get("/web*", func(c fiber.Ctx) error {
    return c.SendFile("dist/index.html")
})
```

### Pre Fork
Jika kita menjalankan beberapa aplikasi Fiber secara bersamaan, maka kita harus menggunakan port yang berbeda untuk setiap aplikasi. Pre Fork adalah fitur untuk menjalankan aplikasi Fiber secara bersamaan tanpa harus mengubah port.
```
// V2
fiber.New(fiber.Config{
	Prefork: true,
})
// V3
app.Listen(":3000", fiber.ListenConfig{
    EnablePrefork: true,
})
```

### Error Handling
```
fiber.New(fiber.Config{
	ErrorHandler: func(ctx fiber.Ctx, err error) error {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString("Error: " + err.Error())
	},
})
```

### Template
Mustache adalah templating engine yang digunakan untuk menghasilkan HTML dinamis.

### Middleware
ctx.Next()

### Middleware Lainnya
Fiber menyediakan banyak sekali middleware yang dapat digunakan untuk mempermudah pengembangan aplikasi seperti RequestId, BasicAuth, FileSystem, ETag, dan banyak lagi.
> https://docs.gofiber.io/category/-middleware/

### HTTP Client
fiber.AcquireClient()
``` go
client := client.New().SetBaseURL("https://example.com/")

resp, err := client.Get("/")
if err != nil {
	panic(err)
}
defer resp.Close()
```
