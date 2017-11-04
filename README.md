# txtTransformer

This Golang package is a personal **micro**library to help developing scripts that transforms large numbers of text files. It's being built to make programatic changes to thousands of html files.

## File functions

Functions to read, write, search and filter files and paths.

### FindFiles

Finds all files recursively, ignoring folders.

```go
filesInsideFoo := txttransformer.FindFiles("testdata/foo")
```

### FilterPaths

Filters paths by extension. Use without the dot, like html or txt.

```go
listTxtPaths := txttransformer.FilterPaths(filesInsideFoo, "txt")
```

### PrintListPaths

Prints the list of files.

```go
txttransformer.PrintListPaths(listTxtPaths)
```

### CreatePathFolder

Checks if the folder for the current path exists and creates it if not.

```go
txttransformer.CreatePathFolder("testdata/foo/bar/file.txt")
```

### ReadTextFile

Reads a file into a string.

```go
myTxtFile := txttransformer.ReadTextFile("testdata/file.txt")
```

### WriteTextFile

Writes a string as a file.

```go
txttransformer.WriteTextFile("testdata/foo/bar/file.txt", myTxtFile)
```

## Strings (files as strings)

### ReplaceAllText

Replaces text in multiline strings.

```go
changedText := txttransformer.ReplaceAllText(originalText, "is", "é")
```

### ReplaceMatchedText

Replaces text that matches the regular expression by new text.

```go
changedText := txttransformer.ReplaceMatchedText(originalText, `do\!`, "Xooer")
```

### ReplaceMatchedTextFunction

Replaces text caught in regularExp by the result of the function.

```go
changedText := txttransformer.ReplaceMatchedTextFunction( originalText, `do\!$`, func(inputExp string) string {
    if (inputExp == "do!") {
        return "Work"
    }
    return "Sleep"
})
```
