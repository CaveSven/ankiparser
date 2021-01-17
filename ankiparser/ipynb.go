package main

import (
	"encoding/json"
	"log"
	"strings"
)

type Cells struct {
	Cells []Cell `json:"cells"`
}

type Cell struct {
	CellType string   `json:"cell_type"`
	Source   []string `json:"source"`
}

func readQAndAFromIpynb(filename string) []string {
	data := ReadFile(filename)
	var cells Cells
	err := json.Unmarshal(data, &cells)

	if err != nil {
		log.Fatal(err)
	}

	var text []string
	for _, cell := range cells.Cells {
		if cell.CellType == "markdown" {
			text = parseQAndAFromCell(cell, text)
		}
	}

	return text
}

// Cell looks like
// [
//   "**Q:** What are sights in Berlin?\n",
//   "\n",
//   ">When visiting Berlin, several sights are interesting to visit.\n",
//   ">\n",
//   "> Things to look out for are\n",
//   ">\n",
//   ">    * The Brandenburg Gate.\n",
//   ">    * SO36."
// ]
func parseQAndAFromCell(cell Cell, text []string) []string {
	var multiLineAnswer []string

	for _, line := range cell.Source {
		if lineIsQuestion(line) {
			text = append(text, line)
		} else if isPartOfMultiLineAnswer(line) {
			multiLineAnswer = append(multiLineAnswer, line)
		} else if lineMarksEndOfMultiLineAnswer(line) && len(multiLineAnswer) != 0 {
			text = append(text, flattenMultiLineAnswer(multiLineAnswer))
			multiLineAnswer = []string{}
		}
	}

	if len(multiLineAnswer) != 0 {
		text = append(text, flattenMultiLineAnswer(multiLineAnswer))
	}
	return text
}

func isPartOfMultiLineAnswer(line string) bool {
	return strings.HasPrefix(line, MultiAPrefix)
}

func lineMarksEndOfMultiLineAnswer(line string) bool {
	return line == "\n"
}

func flattenMultiLineAnswer(multiLineAnswer []string) string {
	var answer = APrefix
	for _, multiLinePart := range multiLineAnswer {
		answer += " " + afterPrefix(multiLinePart, MultiAPrefix)
	}
	return answer
}
