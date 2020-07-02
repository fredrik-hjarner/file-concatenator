package paths

import (
	"github.com/fredrik-hjarner/file-concatinator/config"
	"github.com/thoas/go-funk"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetAllFilePaths() []string {
	var result = make([]string, 0)
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		result = append(result, path)
		return nil
	}

	err := filepath.Walk(".", walkFunc)
	if err != nil {
		log.Println(err)
	}
	return result
}

func GetWhitelistedPaths(paths []string) []string {
	hasSuffix := func(test string) func(string) bool {
		return func(suffix string) bool {
			return strings.HasSuffix(test, suffix)
		}
	}

	return funk.FilterString(paths, func(s string) bool {
		_, found := funk.FindString(config.SuffixWhitelist, hasSuffix(s))
		return found
	})
}
