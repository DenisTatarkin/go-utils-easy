package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const FILE_WITH_LIST = "remove_ignore.txt"

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ignore := make(map[string]bool)
	ignoreFile, err := os.Open(FILE_WITH_LIST)
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	input := bufio.NewScanner(ignoreFile)
	for input.Scan() {
		ignore[input.Text()] = true
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	for _, file := range files {
		if ignore[file.Name()] {
			continue
		}

		err := os.Remove(file.Name())
		if err != nil {
			os.RemoveAll(file.Name())
		}
	}

}
