package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Input interface {
	ReadLine() (string, error)
}

type ReaderInput struct {
	reader *bufio.Reader
}

// NewReaderInput creates ReaderInput instance with specified io.Reader
func NewReaderInput(rd io.Reader) ReaderInput {
	return ReaderInput{bufio.NewReader(rd)}
}

//ReadLine reads line from bufio.Reader
func (r ReaderInput) ReadLine() (string, error) {
	return (*r.reader).ReadString('\n')
}

type Output interface {
	WriteLine(line string) error
	Flush() error
}

type StdOutput struct{}

// WriteLine write line to stdout.
// Always return nil
func (o StdOutput) WriteLine(line string) error {
	fmt.Print(line)
	return nil
}

// Flush always return nil
func (o StdOutput) Flush() error {
	return nil
}

type FileOutput struct {
	file   *io.Writer
	buffer *bytes.Buffer
}

func (o FileOutput) WriteLine(line string) error {
	if o.buffer == nil {
		o.buffer = bytes.NewBufferString(line)
	} else {
		o.buffer.WriteString(line)
	}
	return nil
}

type Processor interface {
	Process(i *Input, o *Output, data map[string]string)
}

type SimpleProcessor struct {
	prefix  string
	postfix string
}

func (p SimpleProcessor) Process(i *Input, o *Output, data map[string]string) {
	for {
		inputLine, err := (*i).ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		resultLine := p.processLine(inputLine, data)

		(*o).WriteLine(resultLine)
	}
	(*o).Flush()
}

func (p SimpleProcessor) processLine(line string, data map[string]string) string {
	resultLine := line
	for k, v := range data {
		key := p.getKey(k)
		resultLine = strings.Replace(resultLine, key, v, -1)
	}
	return resultLine
}

func (p SimpleProcessor) getKey(s string) string {
	return p.prefix + s + p.postfix
}
