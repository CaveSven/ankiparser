package main

import (
	"testing"
)

func TestCardString(t *testing.T) {
	expectedText := "AnkiCard\n[Q] What is this?\n[A] This is me!"
	actualText := (&AnkiCard{"What is this?", "This is me!"}).String()

	if actualText != expectedText {
		t.Errorf("Expected %v but got %v", expectedText, actualText)
	}
}

func TestCardExport(t *testing.T) {
	expectedText := "What is this?\tThis is me!"
	actualText := (&AnkiCard{"What is this?", "This is me!"}).Export()

	if actualText != expectedText {
		t.Errorf("Expected %v but got %v", expectedText, actualText)
	}
}
