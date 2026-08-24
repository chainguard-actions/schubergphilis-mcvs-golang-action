package main

import "fmt"

// Version is set at build time.
var Version = "dev"

func main() {
	fmt.Printf("Hello from testapp version %s\n", Version)
}
