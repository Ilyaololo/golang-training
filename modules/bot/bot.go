package bot

import "fmt"

type Bot interface {
	GetGreeting() string
}

type EnglishBot struct{}

type SpanishBot struct{}

func (EnglishBot) GetGreeting() string {
	return "H1 There!"
}

func (SpanishBot) GetGreeting() string {
	return "Hola"
}

func PrintGreeting(b Bot) {
	fmt.Println(b.GetGreeting())
}

func init() {
	eb := EnglishBot{}
	sp := SpanishBot{}

	PrintGreeting(eb)
	PrintGreeting(sp)
}
