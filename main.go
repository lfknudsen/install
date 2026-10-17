package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hymkor/trash-go"
)

func findBinary() string {
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		if info.Mode().IsRegular() {
			ext := filepath.Ext(info.Name())
			if ext == ".exe" || ext == "" {
				prefix, _ := strings.CutSuffix(info.Name(), ext)
				base := filepath.Base(prefix)
				dir, _ := os.Getwd()
				dir = filepath.Base(dir)
				if base == dir {
					return info.Name()
				}
			}
		}
	}
	return ""
}

func install(filename string, installLocation string) {
	filename = filepath.Base(filename)
	dst := filepath.Join(installLocation, filename)
	err := trash.Throw(dst)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}
	err = os.Rename(filename, dst)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}

func main() {
	loc, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	installLocation := filepath.Dir(loc)
	if len(os.Args) < 2 {
		if bin := findBinary(); bin != "" {
			install(bin, installLocation)
		} else {
			fmt.Println("Could not locate a suitable binary. Please specify its name:\n" +
				"$ install <filename(s)...>")
			os.Exit(1)
		}
	} else {
		for _, filename := range os.Args[1:] {
			install(filename, installLocation)
		}
	}
}
