package main

import "fmt"

// Version is set at build time via ldflags.
var Version = "dev"

func main() {
	fmt.Printf("Hello from mcvs-golang-action test app version %s\n", Version)
}
