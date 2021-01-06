package main

import (
	"flag"
	"log"
	"os"
)

var filenameFlag string

func init() {
	flag.StringVar(&filenameFlag, "filename", "", "Path to notes file")
	log.SetPrefix("ankiparser: ")
	log.SetFlags(0)

}

func parseFlag() string {
	flag.Parse()
	if filenameFlag == "" {
		log.Fatal("No filename provided.")
	} else if _, err := os.Stat(filenameFlag); err != nil {
		log.Fatalf("File [%v] does not exist", filenameFlag)
	}

	return filenameFlag
}

func main() {
	filename := parseFlag()
	log.Printf("Parsing file [%v]", filename)
}
