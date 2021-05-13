package main

func main() {
	cards := newDeckFromFile("deck.txt")
	cards.print()
	cards.shuffle()
	cards.shuffle()
	cards.shuffle()
	cards.shuffle()
	cards.print()
}
