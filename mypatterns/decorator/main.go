package main

import "fmt"

type Reader interface {
	Read() string
}

type FileReader struct{}

func (f *FileReader) Read() string {
	return "data from file"
}

type BufferedReader struct {
	reader Reader
}

func (b *BufferedReader) Read() string {
	data := b.reader.Read()
	return "buffered " + data
}

type DecryptReader struct {
	reader Reader
}

func (d *DecryptReader) Read() string {
	data := d.reader.Read()
	return "decrypt " + data
}

func main() {

	fileReader := &FileReader{}

	// Buffered read
	bufferedReader := &BufferedReader{reader: fileReader}
	fmt.Println(bufferedReader.Read())

	// Buffered and decrypted read
	decryptBufferedReader := &DecryptReader{reader: fileReader}
	fmt.Println(decryptBufferedReader.Read())
}
