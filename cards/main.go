package main

func main() {
	cards := newDeck()
	hand, remainingHand := deal(cards, 3)
	hand.print()
	remainingHand.print()
}
