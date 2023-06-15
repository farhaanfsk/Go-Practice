package main

func main() {
	//cards := newDeck()
	//cards.print()

	//hand, remaining := cards.deal(7)
	//hand.print()
	//remaining.print()

	//fmt.Println(cards.toString())
	//cards.saveToFile("cards")
	cards := readCardsFromFile("card")
	cards.print()

}
