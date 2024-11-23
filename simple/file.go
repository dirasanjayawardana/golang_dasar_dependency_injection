package simple

import "fmt"

// Provider bisa return function untuk melakukan proses cleanup setelah object dibuat
// closure akan otomatis dipanggil dalam proses cleanup oleh Google Wire

type File struct {
	Name string
}

func (f *File) Close() {
	fmt.Println("Close file", f.Name)
}

func NewFile(name string) (*File, func()) {
	file := &File{Name: name}
	return file, func() {
		file.Close()
	}
}
