package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func myUsage() {
	fmt.Printf("Usage: %s [OPTIONS] filename targetdir \n", os.Args[0])
	flag.PrintDefaults()
}

func readArgs() (string, string) {
	flag.Usage = myUsage
	flag.Parse()
	args := os.Args[1:]
	if len(args) != 2 {
		flag.Usage()
		os.Exit(1)
	}
	filename, targetDir := os.Args[1], os.Args[2]
	if _, err := os.Stat(filename); err != nil {
		log.Fatalf("File [%v] does not exist", filename)
	}
	if stat, err := os.Stat(targetDir); err != nil || !stat.IsDir() {
		log.Fatalf("Target directory [%v] does not exist", targetDir)
	}

	return filename, targetDir
}
