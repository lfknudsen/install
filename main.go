package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hymkor/trash-go"
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
		dst := filepath.Join(installLocation, filename)
		err = trash.Throw(dst)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
		}
		err = os.Rename(filename, dst)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}
