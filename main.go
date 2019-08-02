package main

import (
	"bufio"
	"fmt"
	"github.com/christiaan-janssen/org/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var todos []todo

type todo struct {
	title string
	tags  string
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		utils.LogError(err)
		if filepath.Ext(path) == ".org" {
			*files = append(*files, path)
		}
		return nil
	}
}

// ParseOrgFile read a file and returns the todo items
func ParseOrgFile(path string) {
	file, err := os.Open(path)
	utils.LogError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "** TODO") {
			fmt.Println(strings.Fields(line)[2:])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var files []string
	root := "/home/christiaan/org/"

	err := filepath.Walk(root, visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		ParseOrgFile(file)
	}

}
