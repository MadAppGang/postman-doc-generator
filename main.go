package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	appName               = "postman-doc-generator"
	defaultInputFilename  = "postman_collection_tpl.json"
	defaultOutputFilename = "postman_collection.json"
	defaultDirectory      = "."
)

var (
	structNames = flag.String("struct", "", "comma-separated list of struct names; must be set")
	input       = flag.String("input", defaultInputFilename, "file name of postman collection with keywords for replacement they to models")
	output      = flag.String("output", defaultOutputFilename, "output file name")
	dir         = flag.String("dir", defaultDirectory, "directory to find structures")
)

// Usage is a replacement usage function for the flags package
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s [flags] -struct=S\n", appName)
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(appName + ": ")
	flag.Usage = Usage
	flag.Parse()
	if len(*structNames) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	structs := strings.Split(*structNames, ",")

	fmt.Printf("Dir: %s\n", *dir)
	fmt.Printf("Structs for conversion: %+v\n", structs)
	fmt.Printf("Input file name: %s\n", *input)
	fmt.Printf("Output file name: %s\n", *output)
}

// isDirectory returns true if the named file is a directory
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}
