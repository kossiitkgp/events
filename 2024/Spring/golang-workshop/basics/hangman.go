package main

import (
	"fmt"
)

func main() {
	word := "test"

	used := make(map[byte]bool)

	for i := 0; i < 5; i++ {
		fmt.Printf("You have %d chances left.\n", 5-i)
		fmt.Print("Try a character : ")

		var input string
		fmt.Scanln(&input)

		used[input[0]] = true

		found := true

		for _, letter := range word {
			if used[byte(letter)] == true {
				fmt.Printf("%c", letter)
			} else {
				found = false
				fmt.Print("_")
			}
		}
		fmt.Print("\n")

		if found {
			fmt.Println("YOU WON!")
			return
		}
	}
	fmt.Println("YOU LOST")
}
