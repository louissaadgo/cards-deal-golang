package main

func main() {
	cards := newDeck()
	cards.shuffle()
	currentDeck, _ := deal(cards, 13)
	currentDeck.printDeck()
}
