package main

import (
	"flag"
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

	models := generator.GetModels()
	postmanSchema := postman.ParseFile(*flagOutput)
	err := postmanSchema.AddModels(models)
	if err != nil {
		log.Fatalf("fail to get models. %v", err)
	}

	err = postmanSchema.Save(*flagOutput)
	if err != nil {
		log.Fatal(err)
	}
}
