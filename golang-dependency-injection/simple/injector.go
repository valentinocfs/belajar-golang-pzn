//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabasePostgreSQL,
		NewDatabaseMongoDB,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarRepository() *FooBarService {
	wire.Build(
		fooSet,
		barSet,
		NewFooBarService,
	)
	return nil
}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializeHelloService() *HelloService {
	wire.Build(
		helloSet,
		NewHelloService,
	)
	return nil
}

var FooBarSet = wire.NewSet(
	NewFoo,
	NewBar,
)

func InitializeFooBar() *FooBar {
	wire.Build(
		FooBarSet,
		wire.Struct(new(FooBar), "Foo", "Bar"), // * for all fields
	)
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}
var FooBarValueSet = wire.NewSet(
	wire.Value(fooValue),
	wire.Value(barValue),
)

func InitializeFooBarUsingValue() *FooBar {
	wire.Build(
		FooBarValueSet,
		wire.Struct(new(FooBar), "*"),
	)
	return nil
}

func InitializeReader() io.Reader {
	wire.Build(
		wire.InterfaceValue(new(io.Reader), os.Stdin),
	)
	return nil
}

func InitializeConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

func InitializeConnection(name string) (*Connection, func()) {
	wire.Build(
		NewConnection,
		NewFile,
	)
	return nil, nil
}
