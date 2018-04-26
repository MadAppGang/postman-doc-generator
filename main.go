package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

const (
	appName               = "postman-doc-generator"
	defaultOutputFilename = "postman_collection.json"
	defaultDir            = "."
)

var (
	flagStruct = flag.String("struct", "", "comma-separated list of struct names")
	flagFile   = flag.String("file", "", "go filename to be parsed")
	flagDir    = flag.String("dir", defaultDir, "directory to be parsed")
	flagOutput = flag.String("output", defaultOutputFilename, "postman collection filename")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(appName + ": ")
	flag.Parse()

	var structs []string
	if len(*flagStruct) > 0 {
		structs = strings.Split(*flagStruct, ",")
	}

	if *flagFile != "" {
		// ParseFile *flagFile
	} else {
		// ParseDir *flagDir
	}

	fmt.Printf("Dir: %s\n", *flagDir)
	fmt.Printf("Structs for conversion: %+v\n", structs)
	fmt.Printf("Postman filename: %s\n", *flagOutput)
}
