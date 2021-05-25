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

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	f, err := os.Create("list_files.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	output := bufio.NewWriter(f)
	for _, file := range files {
		output.WriteString(file.Name() + "\n")
		err := output.Flush()
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	f.Close()
}
