package blackjack

import (
	"fmt"
	"strings"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	card = strings.ToLower(card)
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
/*
Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:

- If you have a pair of aces you must always split them.
- If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a face card (Jack/Queen/King) or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
- If your cards sum up to a value within the range [17, 20] you should always stand.
- If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
- If your cards sum up to 11 or lower you should always hit.

Commands
- Stand (S)
- Hit (H)
- Split (P)
- Automatically win (W)
*/
func FirstTurn(card1, card2, dealerCard string) string {

	totalValue := ParseCard(card1) + ParseCard(card2)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case (totalValue == 21) && dealerCard != "ace" && ParseCard(dealerCard) != 10:
		return "W"
	case totalValue == 21 && ParseCard(dealerCard) == 10:
		return "S"
	case totalValue <= 20 && totalValue >= 17:
		return "S"
	case (totalValue <= 16 && totalValue >= 12) && ParseCard(dealerCard) < 7: 
		return "S"
	case (totalValue <= 16 && totalValue >= 12) && ParseCard(dealerCard) >= 7: 
		return "H"
	case totalValue < 12 :
		return "H"
	default:
		return "S"
	} 
}
