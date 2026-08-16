package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    value := 0;
	switch (card) {
        case "ace":
        	value = 11
        case "two":
        	value = 2
        case "three":
        	value = 3
    	case "four":
        	value = 4
        case "five":
        	value = 5
        case "six":
        	value = 6
        case "seven":
        	value = 7
        case "eight":
        	value = 8
        case "nine":
        	value = 9
        case "ten", "jack", "queen", "king":
        	value = 10
    }
    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	value1, value2, valueDealer := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
    sum := value1 + value2
    if (sum == 22) {
        // Pair of aces
        return "P"
    } else if (sum == 21) {
        // Blackjack
        if (valueDealer) < 10 {
            return "W"
        }
        return "S" 
    } else if sum >= 17 && sum <= 20 {
        return "S"
    } else if sum >= 12 && sum <= 16 {
        if (valueDealer >= 7) {
            return "H"
        }
        return "S"
    }
    return "H"
}
