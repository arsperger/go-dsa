package main

import "fmt"

func main() {

	var arr [2][3]int

	arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for _, rows := range arr {
		for _, cols := range rows {
			fmt.Printf("[%d]", cols)
		}
		fmt.Println("")
	}

}
