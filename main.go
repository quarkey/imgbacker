package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	src := flag.String("src", "", "Source destinations you'd like to backup. Multiple sources can be seperated by comma ")
	dst := flag.String("dst", "", "Destination folder")
	flag.Parse()

	// argument check
	if *src == "" || *dst == "" {
		log.Fatal("Please specify source and destination arguments.")
	}

	// checking src folders
	fmt.Println("Checking src/dst permissions:")
	err := multistat(*src)
	if err != nil {
		log.Fatal(err)
	}
	// checking dst folders
	fmt.Println("Destination folder:", *dst)
	err = isFolder(*dst)
	if err != nil {
		log.Fatal(err)
	}

	// program goes here:
}

func multistat(paths ...string) error {
	for _, arr := range paths {
		x := strings.Split(arr, ",")
		for _, path := range x {
			file, err := os.Stat(path)
			if err != nil {
				return fmt.Errorf("unable to stat file: %v", err)
			}
			if !file.IsDir() {
				return fmt.Errorf("%s is not a folder", path)
			} else {
				fmt.Println(path, "OK!")
			}
		}
	}
	fmt.Println() //pretty print
	return nil
}

func isFolder(path string) error {
	fileinfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("unable to stat: %v", err)
	}
	if !fileinfo.IsDir() {
		return fmt.Errorf("%s is not a folder", path)
	}
	return nil
}
