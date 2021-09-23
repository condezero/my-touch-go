package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	executableRegex *regexp.Regexp = regexp.MustCompile(`.+\.exe`)
	directoryRegex  *regexp.Regexp = regexp.MustCompile(`[\\]`)
)

func main() {

	args := os.Args[1:]
	if len(os.Args) == 0 {
		fmt.Println("Usage:", os.Args[0], "<FILENAME>")
		return
	}

	for _, fileName := range args {
		if executableRegex.Match([]byte(fileName)) {
			fmt.Println("WARN:")
			fmt.Printf(" .exe could not be used as a file extension for %s\n", fileName)

			continue
		}

		if directoryRegex.Match([]byte(fileName)) {
			splited := strings.Split(fileName, "\\")
			_, err := os.Stat(splited[0])
			if os.IsNotExist(err) {
				err := os.MkdirAll(splited[0], os.ModePerm)
				if err != nil {
					panic(err)
				}
			}

		}
		if _, err := os.Stat(fileName); err == nil {
			fmt.Printf("File %s already exists.\n", fileName)
			continue
		}
		f, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}

}
