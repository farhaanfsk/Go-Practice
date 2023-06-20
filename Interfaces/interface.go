package main

import "fmt"

type bot interface {
	getgreeting() string
}
type bot2 interface {
	getgreeting() string
	getWishings() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb, eb)
	printGreeting(sb, sb)

}

func (eb englishBot) getgreeting() string {
	return "Hi there!"
}

func (sb spanishBot) getgreeting() string {
	return "Hola!"
}

func (eb englishBot) getWishings() string {
	return "Hi there1!"
}

func (sb spanishBot) getWishings() string {
	return "Hola1!"
}

func printGreeting(b bot2, c bot) {
	fmt.Println(c.getgreeting())
	fmt.Println(b.getWishings())
}
