package main

import "fmt"

// Version is overridden at build time via -ldflags "-X 'main.Version=...'".
var Version = "dev"

func main() {
	fmt.Println("version=" + Version)
}
