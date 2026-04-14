package blackjack
import "fmt"
func ParseCard(s string) int {
	switch s {
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

func FirstTurn(card1, card2, dealerCard string) string {
	handValue := ParseCard(card1) + ParseCard(card2)

	
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	
	if handValue == 21 {
		if dealerCard == "ace" || dealerCard == "king" || dealerCard == "queen" || dealerCard == "jack" || dealerCard == "ten" {
			return "S"
		}
		return "W"
	}

	
	if handValue >= 17 && handValue <= 20 {
		return "S"
	}

	
	if handValue >= 12 && handValue <= 16 {
		if dealerCard == "seven" ||
			dealerCard == "eight" ||
			dealerCard == "nine" ||
			dealerCard == "ten" ||
			dealerCard == "jack" ||
			dealerCard == "queen" ||
			dealerCard == "king" ||
			dealerCard == "ace" {
			return "H"
		}
		return "S"
	}

	
	return "H"
}
func main() {
	fmt.Println(ParseCard("ace"))

	fmt.Println(FirstTurn("ace", "ace", "jack"))
	fmt.Println(FirstTurn("ace", "king", "ace"))
	fmt.Println(FirstTurn("five", "queen", "ace"))
}
