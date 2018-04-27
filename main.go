package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/madappgang/postman-doc-generator/models/postman"
)

const (
	appName       = "postman-doc-generator"
	defaultOutput = "postman_collection.json"
	defaultSource = "."
)

var (
	flagStruct = flag.String("struct", "", "comma-separated list of struct names")
	flagSource = flag.String("source", defaultSource, "filename or directory to be parsed")
	flagOutput = flag.String("output", defaultOutput, "postman collection filename")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(appName + ": ")
	flag.Parse()

	var structs []string
	if len(*flagStruct) > 0 {
		structs = strings.Split(*flagStruct, ",")
	}

	generator := NewGenerator(structs)
	generator.ParseSource(*flagSource)

	fmt.Printf("Source: %s\n", *flagSource)
	fmt.Printf("Structs for conversion: %+v\n", structs)
	fmt.Printf("Postman filename: %s\n", *flagOutput)

	models := generator.GetModels()
	fmt.Println("Found models:")
	for _, model := range models {
		fmt.Printf("%s\n", model)
	}

	postmanSchema := postman.ParseFile(inputName)

	var newModels string
	for _, model := range models {
		newModels += model.String()
	}

	postmanSchema.SetModels(newModels)

	err := postmanSchema.Save(outputName)
	if err != nil {
		log.Fatal(err)
	}
}
