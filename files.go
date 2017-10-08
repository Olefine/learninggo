package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	showAllFilesRecursive()
}

func readFile() {
	filePath := "/home/egorodov/geo.csv"
	bs, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	content := string(bs)

	for _, line := range strings.Split(content, "\n") {
		fmt.Println(line)
	}
}

func createFile()  {
	filePath := "/home/egorodov/testfilego.txt"

	file, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	names := []string{"Egor", "Inna", "Oleg"}

	namesJoned := strings.Join(names, "\n") + "\n"

	file.WriteString(namesJoned)
}

func showAllFilesInDir() {
	dir, err := os.Open("/home")
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)

	if err != nil {
		panic(err)
	}

	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}
}

func showAllFilesRecursive() {
	filepath.Walk("/home", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
