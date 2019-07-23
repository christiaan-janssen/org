package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		logError(err)
		if filepath.Ext(path) == ".org" {
			*files = append(*files, path)
		}
		return nil
	}
}

func logError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	/*
		var files []string

		root := "/home/christiaan/org"
		err := filepath.Walk(root, visit(&files))
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			fmt.Println(file)
		}
	*/
	file, err := os.Open("/home/christiaan/org/gtd.org")
	logError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "** TODO") {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
