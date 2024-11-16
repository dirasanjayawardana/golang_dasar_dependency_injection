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