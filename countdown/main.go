package main

import (
	"fmt"
	"math"
	"slices"
)

type Node struct {
	branch      int
	parent      int
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
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter target: ")
	// targetIn, _ := reader.ReadString('\n')

	// target, err := strconv.Atoi(strings.TrimSpace(targetIn))

	// if err != nil {
	// 	fmt.Println("That is not an integer. Error: ")
	// 	return
	// }

	// fmt.Print("Enter array of numbers (1,2,3,...): ")
	// numIn, _ := reader.ReadString('\n')

	// numInArray := strings.Split(numIn, ",")
	// num := make([]int, len(numInArray))

	// for i, line := range numInArray {
	// 	num[i], err = strconv.Atoi(strings.TrimSpace(line))

	// 	if err != nil {
	// 		fmt.Println("That is not an integer. Error: ")
	// 		return
	// 	}
	// }

	target := 494
	//num := []int{50, 100, 75, 25, 5, 7}
	num := []int{100, 5, 7, 75, 25, 50}
	eq := 4 //0 = +, 1 = -, 2 = *, 3 = /
	calc := []Node{}

	length := 0
	oldlength := 0
	closest := Node{total: 0}
	fmt.Printf("Target: %d and numbers is %v\n", target, num)
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
					branch:      loop,
					parent:      i,
					indice:      []int{i},
					calculation: []int{i},
					calc_string: fmt.Sprintf("%d", num[i]),
					total:       num[i],
				}
				calc = append(calc, result)
				if math.Abs(float64(closest.total-target)) > math.Abs(float64(result.total-target)) {
					closest = result
				}
				if result.total == target {
					closest = result
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
					if slices.Contains(selected.indice, j) {
						continue
					}
					// Loop through the arithmatic symbols
					for e := 0; e < eq; e++ {
						newIndice := []int{}
						newCalculation := []int{}
						newIndice = append(newIndice, selected.indice...)
						newCalculation = append(newCalculation, selected.calculation...)
						var result Node
						switch e {
						case 0:
							result = Node{
								branch:      loop,
								parent:      i,
								indice:      append(newIndice, j),
								calculation: append(newCalculation, e, j),
								calc_string: fmt.Sprintf("%s+%d", selected.calc_string, num[j]),
								total:       selected.total + num[j],
							}
						case 1:
							result = Node{
								branch:      loop,
								parent:      i,
								indice:      append(newIndice, j),
								calculation: append(newCalculation, e, j),
								calc_string: fmt.Sprintf("%s-%d", selected.calc_string, num[j]),
								total:       selected.total - num[j],
							}
						case 2:
							result = Node{
								branch:      loop,
								parent:      i,
								indice:      append(newIndice, j),
								calculation: append(newCalculation, e, j),
								calc_string: fmt.Sprintf("%s*%d", selected.calc_string, num[j]),
								total:       selected.total * num[j],
							}
						case 3:
							if selected.total%num[j] != 0 {
								continue
							}
							result = Node{
								branch:      loop,
								parent:      i,
								indice:      append(newIndice, j),
								calculation: append(newCalculation, e, j),
								calc_string: fmt.Sprintf("%s/%d", selected.calc_string, num[j]),
								total:       selected.total / num[j],
							}
						}

						calc = append(calc, result)
						if math.Abs(float64(closest.total-target)) > math.Abs(float64(result.total-target)) {
							closest = result
						}
						if result.total == target {
							closest = result
							break out
						}

					}
				}
			}
		}

		oldlength = length
	}

	fmt.Printf("Closest: %v\n", closest)

	// parent := closest.parent
	// for parent > len(num) {
	// 	parentNode := calc[parent]
	// 	fmt.Printf("Closest: %v = %d\n", parentNode, parent)
	// 	parent = parentNode.parent
	// }
}
