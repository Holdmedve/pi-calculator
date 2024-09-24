package main

import "fmt"
import "time"
// struct PiCalculator {

// }


func increasePrecision(pi float64, n float64, m int) float64 {
	term := 4 / (n * (n+1) * (n+2))
	if m % 2 == 0 {
		return pi - term
	} else {
		return pi + term
	}
}


func CalculatePi(cmdCh <-chan string, resultCh chan<- float64) float64 {
	n := 2.0
	pi := 3.0

	for m := 1; true; m++ {
		if len(cmdCh) == 1 {
			cmd := <-cmdCh
			if cmd == "stop" {
				fmt.Println("Exit command received")
				break
			} else if cmd == "show" {
				resultCh <- pi
			}
		}
		pi = increasePrecision(pi, n, m)
		n += 2.0
		time.Sleep(time.Second * 1)
	}

	resultCh <- pi
	fmt.Println("result sent to result channel")
	return pi
}