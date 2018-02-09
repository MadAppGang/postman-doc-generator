package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	currentDir, err := getCurrentDir()
	if err != nil {
		log.Fatalf("Cannot get current dir. %v", err)
	}

	// get path from arguments or set current dir
	path := flag.String("p", currentDir, "Path to source files")
	flag.Parse()

	// create a new generator
	generator := NewGenerator()
	models, err := generator.ParseAll(*path)
	if err != nil {
		log.Fatalf("Cannot parse files. %v", err)
	}

	// print created models
	fmt.Println(models.String())
}

// getCurrentDir returns the current directory
func getCurrentDir() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(ex), nil
}
