//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

// Injector adalah sebuah function Constructor, namun isinya berupa konfigurasi untuk Google Wire
// Injector tidak akan digunakan oleh program dalam project, hanya untuk melakukan auto generate kode Dependency Injection
// Untuk membuat Injector, harus menambahkan komentar `//go:build wireinject` dan `// +build wireinject`

// Injector akan return Dependecy terakhir yang dibutuhkan dan error jika terjadi error di Provider
// Setiap parameter di Provider akan diinject berdasarkan tipe datanya, akan dicari dari return constructor di wire.Build()
// jika tidak menemukan tipedata yang cocok, maka parameter di Provider akan diinject dari parameter Injector
func InitializedService(isError bool) (*SimpleService, error) {
	// memberi tahu goole wire, function provider mana yang akan digunakan
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	// cukup returnkan nil, karena semua code yg ada di function ini akan ditimpa oleh google wire ketika di generate
	return nil, nil
}

// Multiple Binding, harus menggunakan alias untuk tipe data yang sama
// Karena setiap parameter di Provider akan diinject berdasarkan tipe datanya
// jika ada Provider yg return tipe data sama, maka akan error
func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// Salah
//func InitializedHelloService() *HelloService {
//	wire.Build(NewHelloService, NewSayHelloImpl)
//	return nil
//}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

func InitializedFooBar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
