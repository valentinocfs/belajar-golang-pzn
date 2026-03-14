## Golang Validation

### Validator Package
Package validator yang paling populer digunakan di Golang:
- go-playground/validator

### Validate Struct
Validator Struct akan melakukan cache informasi seperti rules, tags, dan lain-lain. yang berhubungan dengan validation kita. Oleh karena itu, kita tidak perlu melakukan validasi struct berkali-kali.

### Validasi Dua Variable
Untuk melakukan validasi dua variable, kita bisa gunakan method
- Validate.VarWithValue(first, second, "tag/rule")
- Validate.VarWithValueCtx(ctx, first, second, "tag/rule")

### Baked-in Validations (Package Bawaan)
https://pkg.go.dev/github.com/go-playground/validator/v10#readme-special-notes

### Multiple Tag Validation
Untuk melakukan validasi multiple tag, kita bisa gunakan method
- Validate.Var(value, "tag1,tag2,tag3")
- Validate.VarCtx(ctx, value, "tag1,tag2,tag3")

### Tag Parameter
Untuk melakukan validasi dengan parameter, kita bisa gunakan method
- Validate.Var(value, "tag:parameter")
- Validate.VarCtx(ctx, value, "tag:parameter")

### Validasi Struct
```
  type LoginRequest struct {
		Username string `validate:"required,min=5,max=10"`
		Password string `validate:"required,min=8"`
	}
```

### Validasi Error
```
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace()) // nama struct + nama field
			fmt.Println(err.Field()) // nama field
			fmt.Println(err.Tag()) // tag yang gagal
			fmt.Println(err.Param()) // parameter tag
			fmt.Println(err.Kind()) // kind of field
			fmt.Println(err.Type()) // type of field
			fmt.Println(err.Value()) // value of field
		}
```

### Validasi Cross Field
```
	type RegisterRequest struct {
		Username string `validate:"required,min=5,max=10"`
		Password string `validate:"required,min=8"`
		ConfirmPassword string `validate:"required,eqfield=Password"`
	}
```

### Validasi Nested Struct
```
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
		Zip    string `validate:"required"`
	}

	type User struct {
		Name    string  `validate:"required"`
		Email   string  `validate:"required,email"`
		Address Address `validate:"required"`
	}
```

### Validasi Collection
```
	type Order struct {
		Items []string `validate:"required,dive"`
	}
```

### Validasi Basic Collection
```
	type Order struct {
		Items []string `validate:"dive,required,min=1"`
	}
```

### Validasi Map
```
	type School struct {
		Name string `validate:"required"`
	}

  type User struct {
    Name string `validate:"required"`
    Schools map[string]School `validate:"dive,keys,required,min:2,endkeys,dive"`
  }
```

### Validasi Basic Map
```
	type User struct {
		Id      string         `validate:"required"`
		Name    string         `validate:"required"`
		Wallet  map[string]int `validate:"required,dive,keys,required,min=2,endkeys,required,gt=0"`
	}
```

### Alias Tag
Pada beberapa kasus, kadang kita sering menggunakan tag yang sama, sehingga kita bisa membuat alias tag dengan menggunakan method
- Validate.RegisterAlias(alias, tag)

```
	validate.RegisterAlias("valid", "required,min=3")
```

### Custom Validation
Untuk melakukan custom validation, kita bisa menggunakan method
- Validate.RegisterValidation(tag, validationFunc)

```
	validate.RegisterValidation("valid", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "valid"
	})
```

### Custom Validation Parameter
Untuk mengambil parameter dari custom validation, kita bisa menggunakan method
- field.Param()

```
	func MustValidPin(field validator.FieldLevel) bool {
		length, err := strconv.Atoi(field.Param())
		if err != nil {
			return false
		}
		value := field.Field().String()
		return regexNumber.MatchString(value) && len(value) == length
	}
```

### Or Rule
```
	type Login struct {
		Username string `validate:"required,email|alphanum"`
		Password string `validate:"required,min=5`
	}
```

### Custom Validation Cross Field
Untuk mengambil nilai field lain pada struct, kita bisa menggunakan method
- field.GetStructFieldOK2()

### Struct Level Validation
Untuk meregister struct validation, kita bisa menggunakan method
- Validate.RegisterStructValidation(validationFunc, structType...)

```
	validate.RegisterStructValidation(func(sl validator.StructLevel) {
		
	}, RegisterRequest{})
```