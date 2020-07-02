package main

import (
	"fmt"
	"github.com/fredrik-hjarner/file-concatinator/paths"
	"io/ioutil"
)

func readFile(path string) {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	// fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'

	fmt.Println(str) // print the content as a 'string'
}

func main() {
	result := paths.GetAllFilePaths()
	// fmt.Println(result)
	filtered := paths.GetWhitelistedPaths(result)
	fmt.Println(filtered)
	readFile(filtered[0])
}
