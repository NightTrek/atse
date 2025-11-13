package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Hello")
	os.Exit(0)
	filepath.Join("a", "b")
}
