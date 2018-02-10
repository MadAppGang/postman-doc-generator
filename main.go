package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// get path from arguments or set current dir
	flagPath := flag.String("p", "", "Path to source files")
	flag.Parse()

	if flagPath == nil {
		log.Fatal("Cannot get command line arguments")
	} else if *flagPath == "" {
		currentDir, err := getCurrentDir()
		if err != nil {
			log.Fatalf("Cannot get current dir. %v", err)
		}

		*flagPath = currentDir
	}

	// create a new generator
	generator := NewGenerator()
	models, err := generator.ParseAll(*flagPath)
	if err != nil {
		log.Fatal(err)
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
