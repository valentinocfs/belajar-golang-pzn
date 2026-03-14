package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var validate *validator.Validate = validator.New()

func TestValidation(t *testing.T) {
	if validate == nil {
		t.Error("Validator is nil")
	}
}

func TestValidationField(t *testing.T) {
	user := ""

	err := validate.Var(user, "required")
	fmt.Println(err.Error())
	assert.Error(t, err)
}

func TestValidationTwoVariables(t *testing.T) {
	password := "secret"
	confirmPassword := "secretpassword"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	fmt.Println(err.Error())
	assert.Error(t, err)
}

func TestMultipleTag(t *testing.T) {
	user := "string123456"

	err := validate.Var(user, "required,alphanum")
	assert.NoError(t, err)
}

func TestTagParameter(t *testing.T) {
	user := "string123456"

	err := validate.Var(user, "required,min=5,max=10")
	fmt.Println(err.Error())
	assert.Error(t, err)
}

func TestStructValidation(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,min=5,max=10"`
		Password string `validate:"required,min=8"`
	}

	request := LoginRequest{
		Username: "budi",
		Password: "[PASSWORD]",
	}

	err := validate.Struct(request)
	fmt.Println(err.Error())
	assert.Error(t, err)
}

func TestStructValidationWithErrorMessage(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,min=5,max=10"`
		Password string `validate:"required,min=8"`
	}

	request := LoginRequest{
		Username: "budi",
		Password: "[PASSWORD]",
	}

	err := validate.Struct(request)
	for _, err := range err.(validator.ValidationErrors) {
		fmt.Println(err.Namespace()) // nama struct + nama field
		fmt.Println(err.Field())     // nama field
		fmt.Println(err.Tag())       // tag yang gagal
		fmt.Println(err.Param())     // parameter tag
		fmt.Println(err.Kind())      // kind of field
		fmt.Println(err.Type())      // type of field
		fmt.Println(err.Value())     // value of field
	}
}

func TestStructValidationCrossField(t *testing.T) {
	type RegisterRequest struct {
		Username        string `validate:"required,min=5,max=10"`
		Password        string `validate:"required,min=8"`
		ConfirmPassword string `validate:"required,eqfield=Password"` // eqfield = equal field
	}

	request := RegisterRequest{
		Username:        "johndoe",
		Password:        "[PASSWORD]",
		ConfirmPassword: "[PASSWORD]",
	}

	err := validate.Struct(request)
	assert.NoError(t, err)
}

func TestNestedStructValidation(t *testing.T) {
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

	user := User{
		Name:  "budi",
		Email: "budi@gmail.com",
		Address: Address{
			Street: "Jl. Sudirman",
			City:   "Jakarta",
			Zip:    "12345",
		},
	}

	err := validate.Struct(user)
	assert.NoError(t, err)
}

func TestCollectionValidation(t *testing.T) {
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
		Zip    string `validate:"required"`
	}

	type User struct {
		Name      string    `validate:"required"`
		Email     string    `validate:"required,email"`
		Addresses []Address `validate:"required,dive"`
	}

	user := User{
		Name:  "budi",
		Email: "budi@gmail.com",
		Addresses: []Address{
			{
				Street: "",
				City:   "Jakarta",
				Zip:    "12345",
			},
			{
				Street: "",
				City:   "Jakarta",
				Zip:    "12345",
			},
		},
	}

	err := validate.Struct(user)
	fmt.Println(err.Error())
	assert.Error(t, err)
}

func TestBasicCollectionValidation(t *testing.T) {
	type Order struct {
		Items []string `validate:"dive,required,min=1"`
	}

	order := Order{
		Items: []string{"buku", "pensil", "penghapus", ""},
	}

	err := validate.Struct(order)
	fmt.Println(err.Error())
	assert.Error(t, err)
}

func TestMapValidation(t *testing.T) {
	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Name    string            `validate:"required"`
		Schools map[string]School `validate:"required,dive,keys,required,min=2,endkeys"`
	}

	user := User{
		Name: "budi",
		Schools: map[string]School{
			"sd": {
				Name: "",
			},
			"smp": {
				Name: "",
			},
		},
	}

	err := validate.Struct(user)
	fmt.Println(err.Error())
	assert.Error(t, err)
}

func TestBasicMapValidation(t *testing.T) {
	type User struct {
		Id     string         `validate:"required"`
		Name   string         `validate:"required"`
		Wallet map[string]int `validate:"required,dive,keys,required,min=2,endkeys,required,gt=0"`
	}

	user := User{
		Id:   "1",
		Name: "budi",
		Wallet: map[string]int{
			"":        -100000,
			"mandiri": 100000,
		},
	}

	err := validate.Struct(user)
	fmt.Println(err.Error())
	assert.Error(t, err)
}

func TestAliasTag(t *testing.T) {
	validate.RegisterAlias("varchar", "required,max=255")

	user := "john"

	err := validate.Var(user, "varchar,min=5")
	fmt.Println(err.Error())
	assert.Error(t, err)
}

func IsValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestCustomValidation(t *testing.T) {
	validate.RegisterValidation("username", IsValidUsername)

	user := "JOHNDOE"

	err := validate.Var(user, "username")
	assert.NoError(t, err)
}

var regexNumber = regexp.MustCompile(`^[0-9]+$`)

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		return false
	}
	value := field.Field().String()
	return regexNumber.MatchString(value) && len(value) == length
}

func TestCustomValidationWithParam(t *testing.T) {
	validate.RegisterValidation("pin", MustValidPin)

	pin := "123456"

	err := validate.Var(pin, "number,pin=6")
	assert.NoError(t, err)
}

func TestOrRule(t *testing.T) {
	type Login struct {
		Username string `validate:"required,email|alphanum"`
		Password string `validate:"required,min=5"`
	}

	login := Login{
		Username: "budi123456", // harus email atau alphanum
		Password: "[PASSWORD]",
	}

	err := validate.Struct(login)
	assert.NoError(t, err)
}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		return false
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCustomValidationCrossField(t *testing.T) {
	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	type RegisterRequest struct {
		Username string `validate:"required,field_equals_ignore_case=Email|field_equals_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
	}

	request := RegisterRequest{
		Username: "1234567890", // harus sama dengan email atau phone
		Email:    "johndoe@gmail.com",
		Phone:    "1234567890",
	}

	err := validate.Struct(request)
	assert.NoError(t, err)
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegisterSucces(level validator.StructLevel) {
	request := level.Current().Interface().(RegisterRequest)

	if request.Username == request.Email || request.Username == request.Phone {
		// success
	} else {
		// faile
		level.ReportError(request.Username, "Username", "Username", "username", "")
	}
}

func TestStructValidationWithStructLevel(t *testing.T) {
	validate.RegisterStructValidation(MustValidRegisterSucces, RegisterRequest{})

	request := RegisterRequest{
		Username: "johndoe@gmail.com", // harus sama dengan email atau phone
		Email:    "johndoe@gmail.com",
		Phone:    "1234567890",
		Password: "[PASSWORD]",
	}

	err := validate.Struct(request)
	assert.NoError(t, err)
}
