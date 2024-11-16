package main

import "fmt"

func main() {
	numbers := []int{
		1, 2,
	}

	target := 5

	isExist := checkSumOfNumberToTargetNumber(numbers, target)

	fmt.Println(isExist)

}

func checkSumOfNumberToTargetNumber(numbers []int, target int) bool {
	uniqueNumbers := make(map[int]bool)
	var total int

	for i := 0; i < len(numbers); i++ {
		_, ok := uniqueNumbers[numbers[i]]
		if !ok {
			uniqueNumbers[numbers[i]] = true
		}
	}

	for k, _ := range uniqueNumbers {
		for k2, _ := range uniqueNumbers {
			if k != k2 {
				total = k + k2
			}

			if total == target {
				return true
			}

			total = 0
		}
	}

	return false
}
