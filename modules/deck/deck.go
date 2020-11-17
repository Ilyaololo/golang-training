package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []string

func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d Deck) ToString() string {
	return strings.Join(d, ",")
}

func (d Deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())

	for i := range d {
		np := rand.Intn(len(d) - 1)

		d[i], d[np] = d[np], d[i]
	}
}

func NewDeck() Deck {
	cards := Deck{}

	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func NewDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), ",")
}

func init() {
	d := NewDeck()
	d.Print()
}
