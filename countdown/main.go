package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
	indice      []int
	calculation []int
	calc_string string
	total       int
}

/*
	Give an array of numbers and a target, use the array of numbers and arithmatic
	operations +, - and * only, try to real the target if possible

	If not possible, try to find the closest
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter target: ")
	targetIn, _ := reader.ReadString('\n')

	target, err := strconv.Atoi(strings.TrimSpace(targetIn))

	if err != nil {
		fmt.Println("That is not an integer. Error: ")
		return
	}

	fmt.Print("Enter array of numbers (1,2,3,...): ")
	numIn, _ := reader.ReadString('\n')

	numInArray := strings.Split(numIn, ",")
	num := make([]int, len(numInArray))

	for i, line := range numInArray {
		num[i], err = strconv.Atoi(strings.TrimSpace(line))

		if err != nil {
			fmt.Println("That is not an integer. Error: ")
			return
		}
	}

	eq := []int{0, 1, 2, 3} //0 = +, 1 = -, 2 = *, 3 = /
	calc := []Node{}

	length := 0
	oldlength := 0
	closest := Node{total: 0}
out:
	// For each level of the tree
	// First branch = base number n, n+1, n+2...
	// Second branch = n (equation) n+1, n (equation) n+2
	for loop := 0; loop < len(num); loop++ {
		length = len(calc)

		if loop == 0 {
			// Initialization
			for i := 0; i < len(num); i++ {
				result := Node{
					indice:      []int{i},
					calculation: []int{num[i]},
					calc_string: fmt.Sprintf("%d", num[i]),
					total:       num[i],
				}
				calc = append(calc, result)
				if math.Abs(float64(closest.total-target)) > math.Abs(float64(result.total-target)) {
					closest = result
				}
				if result.total == target {
					break out
				}
			}
		} else {
			// for prev len of tree ... new len of tree
			// to calculate for new level of branches
			for i := oldlength; i < length; i++ {
				selected := calc[i]
				// Loop through each number
				for j := 0; j < len(num); j++ {
					// Ignore the numbers already in the equation
					if !slices.Contains(selected.indice, j) {
						// Loop through the arithmatic symbols
					equation:
						for e := 0; e < len(eq); e++ {
							var result Node
							switch eq[e] {
							case 0:
								result = Node{
									indice:      append(selected.indice, j),
									calculation: append(selected.calculation, e, num[j]),
									calc_string: fmt.Sprintf("%s+%d", selected.calc_string, num[j]),
									total:       selected.total + num[j],
								}
							case 1:
								result = Node{
									indice:      append(selected.indice, j),
									calculation: append(selected.calculation, e, num[j]),
									calc_string: fmt.Sprintf("%s-%d", selected.calc_string, num[j]),
									total:       selected.total - num[j],
								}
							case 2:
								result = Node{
									indice:      append(selected.indice, j),
									calculation: append(selected.calculation, e, num[j]),
									calc_string: fmt.Sprintf("%s*%d", selected.calc_string, num[j]),
									total:       selected.total * num[j],
								}
							case 3:
								if selected.total%num[j] == 0 {
									result = Node{
										indice:      append(selected.indice, j),
										calculation: append(selected.calculation, e, num[j]),
										calc_string: fmt.Sprintf("%s/%d", selected.calc_string, num[j]),
										total:       selected.total / num[j],
									}
								} else {
									break equation
								}
							}

							calc = append(calc, result)
							if math.Abs(float64(closest.total-target)) > math.Abs(float64(result.total-target)) {
								closest = result
							}
							if result.total == target {
								break out
							}
						}
					}
				}
			}
		}
		oldlength = length
	}

	for i, e := range calc {
		fmt.Printf("%d %s=%d\n", i, e.calc_string, e.total)
	}

	fmt.Printf("Closest: %s=%d\n", closest.calc_string, closest.total)
}
