package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getFileList() []string {
	var result []string
	for _, arg := range flag.Args() {
		if strings.ContainsAny(arg, `*?[\`) {
			// Wildcard, use filepath.Glob to expand the wildcard
			l, _ := filepath.Glob(arg)
			result = append(result, l...)
		} else {
			// Filename, use it as it is
			result = append(result, arg)
		}
	}
	return result
}

func checkFileList(fileList []string) {
	for _, filename := range fileList {
		stat, err := os.Stat(filename)
		if os.IsNotExist(err) {
			fmt.Printf("The file '%s' does not exist!\n", filename)
			os.Exit(1)
		}
		if stat.IsDir() {
			fmt.Printf("The file '%s' is not a regular file!\n", filename)
			os.Exit(1)
		}
	}
}

func main() {
	fmt.Println("Use hashing method:", optionMethod)
	fmt.Println("Keep extensions:", optionExtension)

	fileList := getFileList()
	checkFileList(fileList)
	maxNameLen := 0 // Used for a pretty formatted output
	for _, filename := range fileList {
		if len(filename) > maxNameLen {
			maxNameLen = len(filename)
		}
	}

	for _, filename := range fileList {
		newFilename, err := hashFile(filename)
		if err != nil {
			fmt.Println("An error occurred:", err)
			os.Exit(1)
		}

		if optionExtension { // Add the extension from the old name to the new name
			ext := filepath.Ext(filename)
			newFilename += ext
		}

		// First we build the format string f.
		// The result will look like "%15s -> %s\n"
		f := "%" + strconv.Itoa(maxNameLen) + "s -> %s\n"
		fmt.Printf(f, filename, newFilename)

		if !optionDry {
			err = os.Rename(filename, newFilename)
			if err != nil {
				fmt.Println("An error occurred renaming", filename, ":", err)
				os.Exit(1)
			}
		}
	}
}
