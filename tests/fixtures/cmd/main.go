package main

import "fmt"

// Version is set at build time via ldflags.
var Version = "dev"

func main() {
	fmt.Printf("hello version %s\n", Version)
}
