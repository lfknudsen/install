package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: install <filename(s)...>")
		os.Exit(1)
	}
	loc, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	installLocation := filepath.Dir(loc)

	for _, filename := range os.Args[1:] {
		filename = filepath.Base(filename)
		err = os.Rename(filename, filepath.Join(installLocation, filename))
		if err != nil {
			fmt.Println(err)
		}
	}
}
