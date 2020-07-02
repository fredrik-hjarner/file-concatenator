package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fredrik-hjarner/file-concatinator/paths"
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

	// count the number of newlines
	// lines := strings.Split(result, "\n")
	// fmt.Printf("lines: %v\n", len(lines))
	return result
	// check(writeErr)
	// 	_, writeErr = f.WriteString("\f")
}

func splitIntoPages(fileString string) string {
	// pages := make([][]string, 0)
	return fileString
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
