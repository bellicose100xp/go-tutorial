package main

import (
)


func main() {
	cards := newDeckFromFile("cards.txt")
	cards.shuffle()
	cards.print()
}
