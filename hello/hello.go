package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

// TO run in the go-workspaces-tutorial directory:
// $ go run ./hello/hello.go
// or
// $ go run github.com/clebersonp/hello

func main() {
	fmt.Println(stringutil.Reverse("Hello"))
	fmt.Println(stringutil.ToUpper("hello"))
}