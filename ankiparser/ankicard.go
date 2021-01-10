package main

import "fmt"

type AnkiCard struct {
	q string
	a string
}

func (a *AnkiCard) String() string {
	return fmt.Sprintf("AnkiCard\n[Q] %v\n[A] %v", a.q, a.a)
}

func (a *AnkiCard) Export() string {
	return fmt.Sprintf("%v\t%v", a.q, a.a)
}
