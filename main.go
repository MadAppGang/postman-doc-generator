package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	appName               = "postman-doc-generator"
	defaultInputFileName  = "postman_collection_tpl.json"
	defaultOutputFileName = "postman_collection.json"
)

var (
	structNames = flag.String("struct", "", "comma-separated list of struct names; must be set")
	input       = flag.String("input", defaultInputFileName, "file name of postman collection with keywords for replacement they to models; default srcdir/"+defaultInputFileName)
	output      = flag.String("output", defaultOutputFileName, "output file name; default srcdir/"+defaultOutputFileName)
)

// Usage is a replacement usage function for the flags package
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s [flags] -struct S [directory]\n", appName)
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
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	generator := NewGenerator(structs)

	var dir string
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
		generator.ParseDir(dir)
	} else {
		dir = filepath.Dir(args[0])
		generator.ParseFiles(args)
	}

	inputName := *input
	outputName := *output

	fmt.Printf("Dir: %s\n", dir)
	fmt.Printf("Structs for conversion: %+v\n", structs)
	fmt.Printf("Input file name: %s\n", inputName)
	fmt.Printf("Output file name: %s\n", outputName)

	models := generator.GetModels()
	fmt.Println("Found models:")
	for _, model := range models {
		fmt.Printf("%s\n", model)
	}
}

// isDirectory returns true if the named file is a directory
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}
