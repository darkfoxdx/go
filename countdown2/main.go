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
	calc_string string
	total       int
}

func containsSame(first []int, second []int) bool {
	for _, f := range first {
		for _, s := range second {
			if f == s {
				return true
			}
		}
	}
	return false
}

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

	// target := 365
	// num := []int{25, 100, 75, 50, 9, 7}

	fmt.Printf("Target: %d and numbers is %v\n", target, num)

	slices.Sort(num)
	slices.Reverse(num)

	calc := []Node{}
	closest := Node{total: 0}
	// Level 1
	for i := 0; i < len(num); i++ {
		result := Node{
			indice:      []int{i},
			calc_string: fmt.Sprintf("%d", num[i]),
			total:       num[i],
		}
		calc = append(calc, result)
		if math.Abs(float64(closest.total-target)) > math.Abs(float64(result.total-target)) {
			closest = result
		}
		if result.total == target {
			closest = result
			break
		}
	}

	// Level all
out:
	for loop := 0; loop < len(num)-1; loop++ {

		length := len(calc)
		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				first := calc[i]
				second := calc[j]

				if i == j || containsSame(first.indice, second.indice) {
					continue
				}

				newIncides := []int{}
				newIncides = append(newIncides, first.indice...)
				newIncides = append(newIncides, second.indice...)
				for e := 0; e < 4; e++ {
					var result Node
					switch e {
					case 0:
						result = Node{
							indice:      newIncides,
							calc_string: fmt.Sprintf("(%s+%s)", first.calc_string, second.calc_string),
							total:       first.total + second.total,
						}
					case 1:
						result = Node{
							indice:      newIncides,
							calc_string: fmt.Sprintf("(%s-%s)", first.calc_string, second.calc_string),
							total:       first.total - second.total,
						}
					case 2:
						result = Node{
							indice:      newIncides,
							calc_string: fmt.Sprintf("%s*%s", first.calc_string, second.calc_string),
							total:       first.total * second.total,
						}
					case 3:
						if second.total == 0 || first.total%second.total != 0 {
							continue
						}
						result = Node{
							indice:      newIncides,
							calc_string: fmt.Sprintf("%s/%s", first.calc_string, second.calc_string),
							total:       first.total / second.total,
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

	fmt.Printf("Closest: %v\n", closest)
}
