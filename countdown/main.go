package main

import "fmt"

func main() {
	//	target := 8
	num := []int{2, 3, 6}
	eq := []int{0, 1, 2} //0 = +, 1 = -, 2 =
	//calc := [][]int{}

	//	found := false
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(eq); j++ {
			var result int
			switch j {
			case 0:
				result = num[i] + num[i+1]
			case 1:
				result = num[i] - num[i+1]
			case 2:
				result = num[i] * num[i+1]
			}
			// calc[i*j+j] = []int{num[i], j, num[i+1], result}
			fmt.Printf("%d %d, %d, %d, %d\n", i*len(num)+j, num[i], j, num[i+1], result)
		}
	}

}
