package main

import (
	"bufio"
	"fmt"
	"os"
)

func fileFeatures() {
	// Multiple methods to deal with the files
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	data, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("File Name", data.Name())
	fmt.Println("File Size", data.Size())
	fmt.Println("Last Modification Time", data.ModTime())
	fmt.Println("Is Folder", data.IsDir())
}

func fileRead() {
	// reading a File

	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	data, err := f.Stat()
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// We store the file data in buffer Buffer is nothing but a place in the memory where you save data
	// Buffer is just an array of bytes
	buff := make([]byte, data.Size())
	d, err := f.Read(buff)
	if err != nil {
		panic(err)
	}
	result := fmt.Sprintf("%d\n%d", d, buff)
	fmt.Println(result)

	// Using os.ReadFile() donot use this one it will make your file load in the memory all at once
	another, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(another))
}

func ReadFolders() {

	// We can even read folders
	dir, err := os.Open("..")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	fileinfo, err := dir.ReadDir(-1)
	if err != nil {
		panic(err)
	}
	for _, fi := range fileinfo {
		fmt.Println(fi.Name())
	}
}

func writeFile() {
	f, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// Append mode
	f.WriteString("hi golang")
	// Append mode
	bytes := []byte("Nice")
	f.Write(bytes)
}

func StreamingReadWrite() {
	// read and write to another file using streaming
	source, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer source.Close()
	dest, err := os.Create("exampledest.txt")
	if err != nil {
		panic(err)
	}
	defer dest.Close()
	// wanna use streaming use built in library bufio
	// Create a reader first
	reader := bufio.NewReader(source)
	writer := bufio.NewWriter(dest)
	// pipe the flow get the data from reader to writer
	// While loop does not exist in go
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}

	}
	writer.Flush()
	fmt.Println("Success")
}

func FileDel() {
	err := os.Remove("exampledest.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File Deleted")
}

func main() {

	// fileFeatures()
	// fileRead()
	// ReadFolders()
	// writeFile()
	// StreamingReadWrite()
	FileDel()

}
