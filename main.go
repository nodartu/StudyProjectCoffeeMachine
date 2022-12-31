package main

import (
	"fmt"
)

func remaining(s []int) {
	//Status of machine

	fmt.Println("The coffee machine has:")
	fmt.Println(s[0], "ml of water")
	fmt.Println(s[1], "ml of milk")
	fmt.Println(s[2], "g of coffee beans")
	fmt.Println(s[3], "disposable cups")
	fmt.Printf("$%d of money\n", s[4])
	fmt.Println()
}
func check(s []int, water, milk, beans, cups, money int) {
	if s[0] < water {
		fmt.Println("Sorry not enough water!")
	} else if s[1] < milk {
		fmt.Println("Sorry, not enough milk!")
	} else if s[2] < beans {
		fmt.Println("Sorry, not enough coffee beans!")
	} else if s[3] < cups {
		fmt.Println("Sorry, not enough cups!")
	} else {
		fmt.Println("I have enough resources, making you a coffee!")
		s[0] -= water
		s[1] -= milk
		s[2] -= beans
		s[3] -= cups
		s[4] += money
	}
	fmt.Println()
}

func espresso(s []int) {
	water, milk, beans, cups, money := 250, 0, 16, 1, 4
	check(s, water, milk, beans, cups, money)
}

func latte(s []int) {
	water, milk, beans, cups, money := 350, 75, 20, 1, 7
	check(s, water, milk, beans, cups, money)
}

func cappuccino(s []int) {
	water, milk, beans, cups, money := 200, 100, 12, 1, 6
	check(s, water, milk, beans, cups, money)
}

func flatwhite(s []int) {
	water, milk, beans, cups, money := 200, 100, 18, 1, 8
	check(s, water, milk, beans, cups, money)
}

func buy(s []int) {
	var selection string
	fmt.Println("What do you want to buy? 1 - Espresso, 2 - Latte, 3 - Cappuccino, 4 - Flat White, back - to main menu:")
	fmt.Scan(&selection)

	switch selection {
	case "1":
		espresso(s)
	case "2":
		latte(s)
	case "3":
		cappuccino(s)
	case "4":
		flatwhite(s)
	case "back":
		break
	default:
		fmt.Println("Invalid selection. Try again.")
		fmt.Println(s)
	}
}

func fill(s []int) {
	var water, milk, beans, cups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&water)
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&milk)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&beans)
	fmt.Println("Write how many disposable cups you want to add:")
	fmt.Scan(&cups)
	fmt.Println()

	s[0] += water
	s[1] += milk
	s[2] += beans
	s[3] += cups
}

func take(s []int) {

	fmt.Printf("I gave you $%d\n", s[4])
	fmt.Println()
	s[4] = 0
}

func actions(s []int) {
	var selection string

	for selection != "exit" {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		fmt.Scan(&selection)
		fmt.Println()

		switch selection {
		case "buy":
			buy(s)
		case "fill":
			fill(s)
		case "take":
			take(s)
		case "remaining":
			remaining(s)
		case "exit":
			break
		default:
			fmt.Println("Invalid selection. Try again.")
		}
	}
}
func main() {
	var s = []int{400, 540, 120, 9, 550}
	actions(s)

}
