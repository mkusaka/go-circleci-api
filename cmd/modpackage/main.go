package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files, err := filepath.Glob("./client/*.go")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("target files:", files)

	for _, file := range files {
		filename := file
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
			return
		}

		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
			return
		}

		if strings.Contains(string(b), "package ") {
			replaced := strings.ReplaceAll(string(b), "package \n", "package client")
			err := ioutil.WriteFile(filename, []byte(replaced), os.ModePerm)
			if err != nil {
				log.Fatal(err)
				return
			}
		}

		fmt.Println("finish update package name for", filename)
	}
}
