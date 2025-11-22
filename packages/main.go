package main

import "hello/hello"

// To create your own package you use go mod init NAME
// convention NAME -> use github repo URL e.g: github.com/shk0Abdullah/go-starter
func main() {
	hello.Hello("coder")
}
