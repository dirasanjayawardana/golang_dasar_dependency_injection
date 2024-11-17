# Dependency Injection
- Dependency Injection merupakan sebuah teknik dimana sebuah object menerima object lain yang dibutuhkan ketika membuat object itu sendiri
- Object yang menerima dependency disebut client, proses mengirim dependencies ke object disebut inject
- Dibahasa pemrograman OOP, untuk membuat object bisa menggunakan Constructor, dan mengirimkan Dependecy yg diperlukan melalui parameter Constructor
- Di Golang harus membuat function untuk membuat object, yang tugasnya mirip seperti Constructor, dan mengirimkan Dependency yang diperlukan melalui parameter function tersebut
- Ketika Project semakin besar, akan sulit harus menentukan urutan object mana yang harus dibuat terlebih dahulu
- Oleh karena itu, proses Dependecy Injection bisa dipermudah dengan memanfaatkan library

## Library Dependecy Injection
- `github.com/google/wire`
- `github.com/uber-go/fx`
- `github.com/golobby/container`

## Google Wire
- Membutuhkan command line wire untuk melakukan auto generate kode Dependency Injection
- `go install github.com/google/wire/cmd/wire@latest`, secar otomatis akan ada file binary di `$GOPATH/bin/wire`
- Untuk memastikan wire telah terinstall, gunakan `wire help`

## Provide 
- Untuk melakukan Dependecy Injection, perlu membuat function Constructor pada struct
- Untuk penamaannya biasanya dengan `NewNamaStructnya()`
- Dalam Google wire, ini disebut Provider

## Injector
- Injector adalah sebuah function Constructor, namun isinya berupa konfigurasi untuk Google Wire
- Injector tidak akan digunakan oleh program dalam project, hanya untuk melakukan auto generate kode Dependency Injection
- Untuk membuat Injector, harus menambahkan komentar `//go:build wireinject` dan `// +build wireinject`

## Proses Dependecy Injection oleh Google Wire
- Setelah membuat Provider dan Injector, jalankan perintah `wire gen namapackage`, dimana namapackage full beserta nama modulenya
- Bisa juga dengan masuk ke direktory package, kemudian cukup jalankan `wire`
- Secara otomatis Google wire akan mencari kode injector di package tersebut
- Lalu membuat file `wire_gen.go` yang isinya adalah function Constructor Injector yang telah terisi otomatis oleh Dependecy
- Sehingga untuk membuat sebuah Object, cukup gunakan function COnstructor yang ada di file `wire_gen.go`

## Error
- Google wire bisa mendeteksi error pada Provider
- Dengan menambahkan return value kedua berupa error pada Injector dan Provider

## Step
- simple/simple.go
- simple/injector.go
- simple/wire_gen.go
- test/simple_service_test.go