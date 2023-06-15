package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	suits := []string{"clubs", "hearts", "spades", "daimonds"}
	values := []string{"king", "queen", "jack", "ace"}
	cards := deck{}
	for _, suit := range suits {
		for _, val := range values {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

func (d deck) deal(n int) (deck, deck) {
	return d[:n], d[n:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(name string) error {
	return ioutil.WriteFile(name, []byte(d.toString()), 0666)
}

func readCardsFromFile(name string) deck {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Print("error:", err)
		return newDeck()
	} else {
		return deck(strings.Split(string(data), ","))
	}
}
