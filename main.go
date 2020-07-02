package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/thoas/go-funk"

	"github.com/fredrik-hjarner/file-concatinator/config"
	"github.com/fredrik-hjarner/file-concatinator/paths"
	"github.com/fredrik-hjarner/file-concatinator/utils"
)

func fileToString(path string) string {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		fmt.Print(err)
		return ""
	}

	return string(b)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func addHeader(path string) string {
	result := "////////////////////////////////////////////////////////////\n"
	result += fmt.Sprintf("File: %v\n", path)
	result += "////////////////////////////////////////////////////////////\n\n"
	result += fileToString(path)
	return result
}

func splitIntoPages(fileString string) string {
	lines := strings.Split(fileString, "\n")
	pages := utils.Chunk(lines, config.MaxLinesPerPage)
	comdensedPages := funk.Map(pages, func(page []string) string {
		return strings.Join(page, "\n") + "\\f\f"
	})
	result := strings.Join(comdensedPages.([]string), "")
	oddNrPages := len(pages)%2 == 1
	if oddNrPages {
		result += "\\f\f"
	}
	return result
}

func main() {
	result := paths.GetAllFilePaths()
	// fmt.Println(result)

	filteredPaths := paths.GetWhitelistedPaths(result)
	// fmt.Println(filteredPaths)

	// fileAsStr := fileToString(filteredPaths[0])
	// fmt.Println(fileAsStr) // print the content as a 'string'

	f, err := os.Create("output.txt")
	check(err)
	defer f.Close()

	// _, writeErr := f.WriteString("Some crap\n")
	// check(writeErr)

	//
	for _, path := range filteredPaths {
		fmt.Println(path)
		withHeader := addHeader(path)
		withFormFeeds := splitIntoPages(withHeader)
		var writeErr error
		_, writeErr = f.WriteString(withFormFeeds)
		check(writeErr)
	}
}
