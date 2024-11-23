package simple

import "fmt"

// Provider bisa return function untuk melakukan proses cleanup setelah object dibuat
// closure akan otomatis dipanggil dalam proses cleanup oleh Google Wire

type Connection struct {
	*File
}

func (c *Connection) Close() {
	fmt.Println("Close connection", c.File.Name)
}

func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{File: file}
	return connection, func() {
		connection.Close()
	}
}
