package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const QPrefix = "**Q:**"
const APrefix = "**A:**"

func init() {
	log.SetPrefix("ankiparser: ")
	log.SetFlags(0)
}

func readQAndAFromFile(filename string) []string {
	log.Printf("Parsing file [%v]", filename)
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		log.Fatal(err)
	}()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	var text []string

	for scanner.Scan() {
		if line := scanner.Text(); strings.Contains(line, APrefix) || strings.Contains(line, QPrefix) {
			text = append(text, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return text
}

func afterPrefix(s string, prefix string) string {
	offset := strings.Index(s, prefix)
	return strings.TrimSpace(s[offset+len(prefix):])
}

func textToCards(text []string) []AnkiCard {
	var cards []AnkiCard

	for i := 0; i < len(text)-1; i++ {
		q, a := text[i], text[i+1]
		if strings.Contains(q, QPrefix) && strings.Contains(a, APrefix) {
			cards = append(cards, AnkiCard{afterPrefix(q, QPrefix), afterPrefix(a, APrefix)})
		}
	}

	return cards
}

func writeCards(cards []AnkiCard, targetDir string) {
	targetFile := filepath.Join(targetDir, "cards.txt")
	f, err := os.Create(targetFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, card := range cards {
		_, err := f.WriteString(card.Export() + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Wrote [%v] cards to file [%s]", len(cards), targetFile)
}

func main() {
	filename, targetDir := readArgs()
	textContent := readQAndAFromFile(filename)
	cards := textToCards(textContent)

	writeCards(cards, targetDir)
}
