package simple

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
}

func NewBar() *Bar {
	return &Bar{}
}

// Struct Provider -> Struct yang bisa dijadikan sebagai Provider (jarang dilakukan)
// Secara otomatis Struct tersebut akan menjadi Provider
// Bisa melakukan dependency injection terhadap field yg ada didalam struct, jika ingin inject semua menggunakan tanda *
type FooBar struct {
	*Foo
	*Bar
}
