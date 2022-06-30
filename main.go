package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var rootPath = flag.String("root", "", "the root path this utility should work in")
var fileExtType = flag.String("ext", "", "a file extension to target")

func main() {
	flag.Parse()

	if *rootPath == "" {
		log.Fatal("must provide root path")
	}
	items, err := os.ReadDir(*rootPath)
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		log.Printf("reading item name: %s", *rootPath+"/"+item.Name())

		if item.IsDir() {
			subItems, err := os.ReadDir(*rootPath + "/" + item.Name())
			if err != nil {
				panic(err)
			}

			for _, subItem := range subItems {
				if !subItem.IsDir() {
					// if a file extension was provided, check for files that match the type in the sub-directory.
					if *fileExtType != "" {
						matches, err := filepath.Match("*."+*fileExtType, subItem.Name())
						if err != nil {
							log.Fatalf("could not match file extension: %s", err.Error())
						}

						// if the file type does not match, skip this one.
						if !matches {
							break
						}
					}
					oldPath := *rootPath + "/" + item.Name() + "/" + subItem.Name()
					newPath := *rootPath + "/" + subItem.Name()

					log.Printf("moving file: %s to path %s", oldPath, newPath)
					os.Rename(oldPath, newPath)
				}
			}
		} else {
			log.Printf("skipping file: %s", item.Name())
		}
	}
}
