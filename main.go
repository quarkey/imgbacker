package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"operation/qlibs/file"
	"os"
	"strings"

	"github.com/quarkey/imgbacker/stack" //https://gist.github.com/bemasher/1777766
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

	// making a stack of all src files
	s := new(stack.Stack)
	for _, dir := range strings.Split(*src, ",") {
		fs, err := file.NewFileSet(dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range fs {
			s.Push(file)
		}
	}
	fmt.Println("Items in stack: ", s.Len())
}

func getmd5(p string) {
	f, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	hash := md5.New()
	_, err = io.Copy(hash, f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x - %s\n", hash.Sum(nil), f.Name())

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
