package blackjack

func ParseCard(card string) int {
	switch card {
        case "ace": return 11
        case "ten", "jack", "queen", "king": return 10
        case "nine": return 9
        case "eight": return 8
        case "seven": return 7
        case "six": return 6
        case "five": return 5
        case "four": return 4
        case "three": return 3
        case "two": return 2
        default: return 0
    }
}

func FirstTurn(card1, card2, dealerCard string) string {
    sumCards := ParseCard(card1) + ParseCard(card2)
	switch {
        case ParseCard(card1) == 11 && ParseCard(card2) == 11: return "P"
        case sumCards == 21 && ParseCard(dealerCard) < 10: return "W"
        case sumCards <= 11: return "H"
        case sumCards <= 16 && ParseCard(dealerCard) >= 7: return "H"
        default: return "S"
    }
}