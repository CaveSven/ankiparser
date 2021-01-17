package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readQAndAFromMd(filename string) []string {
	text := strings.Split(string(ReadFile(filename)), "\n")

	return FilterRows(text)
}

func FilterRows(rows []string) []string {
	var text []string

	for _, row := range rows {
		if lineIsQuestion(row) || strings.Contains(row, APrefix) {
			text = append(text, row)
		}
	}
	return text
}

func lineIsQuestion(row string) bool {
	return strings.Contains(row, QPrefix)
}

func ReadFile(filename string) []byte {
	log.Printf("Parsing file [%v]", filename)
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}
	return data
}
