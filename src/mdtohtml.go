package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

var metadata map[string]string

func FindMetadataDecl(loc int, md []byte) int {
	sepCount := 0
	fmt.Println(len(md))
	for i := loc; i < len(md); i++ {
		if md[i] == '-' {
			sepCount++
		} else {
			sepCount = 0
		}

		if sepCount == 3 {
			return i + 1
		}
	}

	return -1
}

func ParseMetadata(md []byte) []byte {
	metadata = make(map[string]string)
	var beg int = FindMetadataDecl(0, md)
	if beg == -1 {
		return md
	}

	var end int = FindMetadataDecl(beg, md)
	if end == -1 {
		return md
	}

	var otherMd []byte = md[end:]

	readKeyword := true
	var keyword []byte
	var value []byte
	for i := beg + 1; i < end-3; i++ {
		if md[i] == '\n' {
			metadata[string(keyword)] = string(value)
			keyword = nil
			value = nil
			readKeyword = true
		} else if readKeyword == true && md[i] != ' ' {
			// check if trailing ':' has been reached
			if md[i] == ':' {
				readKeyword = false
				i++
				continue
			}

			// read the keyword
			keyword = append(keyword, md[i])
		} else if readKeyword == false && md[i] != '\n' {
			// read the value
			value = append(value, md[i])
		}
	}

	return otherMd
}

func main() {
	extensions := parser.CommonExtensions | parser.Titleblock ^ parser.DefinitionLists
	parser := parser.NewWithExtensions(extensions)

	// check if file name argument is present
	args := os.Args
	if len(args) == 1 {
		os.Stderr.WriteString("Please enter a markdown filename to use!")
		os.Exit(1)
	}

	// read the markdown file contents
	fname := os.Args[1]
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	otherMd := ParseMetadata(content)

	html := markdown.ToHTML(otherMd, parser, nil)
	fmt.Printf(string(html))
}
