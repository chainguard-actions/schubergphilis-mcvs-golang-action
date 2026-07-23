package main

import "fmt"

// Version is set at build time.
var Version = "dev"

func main() {
	fmt.Printf("Hello from test app version %s\n", Version)
}
