package main

import "fmt"


func main() {
	cmdCh := make(chan string, 1)
	resultCh := make(chan float64)

	for true {
		fmt.Println("Type an action:")
		fmt.Println("Calculate Pi (p):")
		fmt.Println("Exit (e):")

		var input string
		var exit bool = false
		fmt.Scanln(&input)
	
		switch(input) {
			case "p":
				go CalculatePi(cmdCh, resultCh)
			case "e":
				cmdCh <- "asd"
				exit = true
		}
		
		if exit {
			fmt.Println(<-resultCh)
			break
		}
	}
}