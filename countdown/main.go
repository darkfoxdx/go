package main

import (
	"fmt"
	"slices"
)

type Node struct {
	indice      []int
	calculation []int
	calc_string string
	total       int
}

func main() {
	//	target := 8
	num := []int{2, 3, 6, 8}
	eq := []int{0, 1, 2} //0 = +, 1 = -, 2 =
	calc := []Node{}

	//	found := false
	for i := 0; i < len(num); i++ {
		result := Node{
			indice:      []int{i},
			calculation: []int{num[i]},
			calc_string: fmt.Sprintf("%d", num[i]),
			total:       num[i],
		}
		calc = append(calc, result)
	}

	length := len(calc)

	for i := 0; i < length; i++ {
		selected := calc[i]
		for j := 0; j < len(num); j++ {
			if !slices.Contains(selected.indice, j) {
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
					}

					calc = append(calc, result)
				}
			}
		}
	}

	oldlength := length
	length = len(calc)

	for i := oldlength; i < length; i++ {
		selected := calc[i]
		for j := 0; j < len(num); j++ {
			if !slices.Contains(selected.indice, j) {
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
					}

					calc = append(calc, result)
				}
			}
		}
	}

	oldlength = length
	length = len(calc)

	for i := oldlength; i < length; i++ {
		selected := calc[i]
		for j := 0; j < len(num); j++ {
			if !slices.Contains(selected.indice, j) {
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
					}

					calc = append(calc, result)
				}
			}
		}
	}

	for i, e := range calc {
		fmt.Printf("%d %s=%d\n", i, e.calc_string, e.total)
	}
}
