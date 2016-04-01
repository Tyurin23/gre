package main

import (
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	inputFile  = kingpin.Flag("input", "Input template file").Short('v').File()
	outputFile = kingpin.Flag("output", "Output file").Short('o').OpenFile(os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0644))
)

func main() {
	kingpin.Parse()

	var input Input
	if *inputFile != nil {
		input = NewReaderInput(*inputFile)
	} else {
		input = NewReaderInput(os.Stdin)
	}

	var output Output
	if *outputFile != nil {

	} else {
		output = &StdOutput{}
	}

	keys := map[string]string{}
	for _, env := range os.Environ() {
		kv := strings.Split(env, "=")
		keys[kv[0]] = kv[1]
	}

	processor := SimpleProcessor{prefix: "{{", postfix: "}}"}
	processor.Process(&input, &output, keys)
}
