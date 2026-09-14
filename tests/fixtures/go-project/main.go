package main

import "fmt"

// Version is set at build time.
var Version = "dev"

func main() {
	fmt.Printf("hello from version %s\n", Version)
}
