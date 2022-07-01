package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		path = filepath.ToSlash(path)
		fmt.Println(path)
		return nil
	})
}
