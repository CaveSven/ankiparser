package main

import (
	"reflect"
	"testing"
)

func TestFindCardsInText(t *testing.T) {
	actualCards := textToCards([]string{"**Q:** What is this?", "**A:** This is me!"})
	expectedCards := []AnkiCard{{"What is this?", "This is me!"}}

	if !reflect.DeepEqual(actualCards, expectedCards) {
		t.Errorf("Expected %v but got %v", expectedCards, actualCards)
	}
}

func TestAfterPrefix(t *testing.T) {
	cases := []struct{ in, prefix, want string }{
		{"{ bla blub **Q:** Test}", QPrefix, "Test}"},
	}

	for _, c := range cases {
		if got := afterPrefix(c.in, c.prefix); got != c.want {
			t.Errorf("AfterPrefix(%q, %q) == %q, want %q", c.in, c.prefix, got, c.want)
		}
	}
}
