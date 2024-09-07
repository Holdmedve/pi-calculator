package main

import "fmt"

func main() {
	for true {
		fmt.Println("Type an action:")
		fmt.Println("Calculate Pi (p):")
		fmt.Println("Exit (e):")

		var input rune
		var exit bool = false
		fmt.Scanf("%c", &input)

		switch(input) {
			case 'p':
				fmt.Println(CalculatePi())
			case 'e':
				exit = true
		}
		
		if exit {
			break
		}
	}
}