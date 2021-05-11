package main

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards.print()

	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, cardSuit := range cardSuits {
		// fmt.Println(cardSuit)
		for _, cardValue := range cardValues {
			// fmt.Println(cardValue + " of " + cardSuit)
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	cards.print()

}

func newCard() string {
	return "Ace of Spades"
}
