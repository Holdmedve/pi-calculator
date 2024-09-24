package main

import "fmt"


func main() {
	cmdCh := make(chan string, 1)
	resultCh := make(chan float64, 0)

	for true {
		fmt.Println("Type an action:")
		fmt.Println("Calculate Pi (p):")
		fmt.Println("Show progress (s):")
		fmt.Println("Exit (e):")

		var input string
		var exit bool = false
		fmt.Scanln(&input)
	
		switch(input) {
			case "p":
				go CalculatePi(cmdCh, resultCh)
			case "e":
				cmdCh <- "stop"
				exit = true
			case "s":
				cmdCh <- "show"
				fmt.Printf("Progress: %.50f\n", <-resultCh)
		}
		
		if exit {
			// if len(resultCh) == 0 {
			// 	fmt.Println("No result to show")
			// 	break
			// }
			fmt.Printf("Result recieved: %.50f\n", <-resultCh)
			break
		}
	}
}