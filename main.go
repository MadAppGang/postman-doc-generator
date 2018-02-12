package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

const defaultPhrase = "$models"

func main() {
	// get path from arguments or set current dir
	flagCollectionPath := flag.String("c", "", "Path to postman collection file. Required")
	flagSourcesPath := flag.String("p", "", "Path to source files. By default uses the current path")
	flagPhrase := flag.String("phrase", defaultPhrase, "The phrase to insert models")
	flag.Parse()

	if *flagCollectionPath == "" {
		log.Fatal("Collection path is required")
	}

	if *flagSourcesPath == "" {
		currentDir, err := getCurrentDir()
		if err != nil {
			log.Fatalf("Cannot get current dir. %v", err)
		}

		*flagSourcesPath = currentDir
	}

	// create a new generator
	generator := NewGenerator()
	err := generator.ParseAll(*flagSourcesPath)
	if err != nil {
		log.Fatal(err)
	}

	generator.Inject(*flagCollectionPath, *flagPhrase)
}

// getCurrentDir returns the current directory
func getCurrentDir() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(ex), nil
}
