package txttransformer

import "strings"
import "regexp"

// ReplaceAllText Replaces text in multiline strings.
func ReplaceAllText(fileText string, searchText string, newText string) string {
	fileText = strings.Replace(fileText, searchText, newText, -1)
	return fileText
}

// ReplaceMatchedText Replaces text that matches the regular expression by new text.
func ReplaceMatchedText(fileText string, regularExp string, newText string) string {
	var re = regexp.MustCompile(regularExp)
	fileText = re.ReplaceAllString(fileText, newText)
	return fileText
}

// ReplaceMatchedTextFunction Replaces text caught in regularExp by the result of the function.
// The function must accept one input param and return a string.
// This funciton is the main purpose of this package, as it will allow many other functions
// (prepend, apend, sub-replace, calculations)
func ReplaceMatchedTextFunction(fileText string, regularExp string, f func(a string) string) string {
	var re = regexp.MustCompile(regularExp)
	fileText = re.ReplaceAllStringFunc(fileText, f)
	return fileText
}
