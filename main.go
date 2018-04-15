package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	src := flag.String("src", "", "Source destinations you'd like to backup. Multiple sources can be seperated with comma ")
	dst := flag.String("dst", "", "Destination folder")
	flag.Parse()

	// argument check
	if *src == "" || *dst == "" {
		log.Fatal("Please specify source and destination arguments.")
	}

	// checking src folders
	err := multistat(*src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// checking dst folders
	err = isFolder(*dst)
	if err != nil {
		log.Println("cannot move on, check the destination folder")
		log.Fatalf("%v", err)
	}
	fmt.Println("destination folder:", *dst)
}

func multistat(paths ...string) error {
	fmt.Println("Checking src permissions:")
	for _, arr := range paths {
		x := strings.Split(arr, ",")
		for _, path := range x {
			fmt.Print(path)
			file, err := os.Stat(path)
			if err != nil {
				return fmt.Errorf(" FAILED! Unable to stat file: %v", err)
			}
			if !file.IsDir() {
				fmt.Println(" FAILED! (not a folder)")
			} else {
				fmt.Println(" OK!")
			}
		}
	}
	return nil
}

func isFolder(path string) error {
	fileinfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("unable to stat: %v", err)
	}
	if !fileinfo.IsDir() {
		return fmt.Errorf("not a folder")
	}
	return nil
}
