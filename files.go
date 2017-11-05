package txttransformer

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// FindFiles Finds all files recursively, ignoring folders.
func FindFiles(searchDir string) []string {
	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() == false {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return fileList
}

// FilterPaths Filters paths by extension. Use without the dot, like html or txt.
func FilterPaths(f []string, ext string) []string {
	r, _ := regexp.Compile(`\.` + ext + `$`)
	var filtered []string
	for _, file := range f {
		if r.MatchString(file) {
			filtered = append(filtered, file)
		}

	}
	return filtered
}

// ReadListPaths Reads a text file with a list of file paths.
func ReadListPaths(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','
	// lineCount := 0

	allRecords, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var allPaths []string
	for _, v := range allRecords {
		allPaths = append(allPaths, v[0])
	}
	return allPaths

}

// PrintListPaths Prints the list of files.
func PrintListPaths(f []string) {
	for _, file := range f {
		fmt.Println(file)
	}
}

// CreatePathFolder Checks if the folder for the current path exists and creates it if not.
func CreatePathFolder(path string) {
	path = filepath.Dir(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
}

// ReadTextFile Reads a file into a string.
func ReadTextFile(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// WriteTextFile Writes a string as a file.
func WriteTextFile(path string, content string) {
	ioutil.WriteFile(path, []byte(content), os.FileMode(0644))
}
