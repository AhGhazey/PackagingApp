package main

import (
	"cmp"
	"fmt"
	"slices"
)

type solution struct {
	currentCombination map[int]int
}

func (s solution) waste(order int) int {
	sum := 0
	for key, value := range s.currentCombination {
		sum += key * value
	}
	return sum - order
}

func (s solution) numberOfPackages() int {
	sum := 0
	for _, value := range s.currentCombination {
		sum += value
	}
	return sum
}

func main() {

	pkg := getBestPackage()
	fmt.Print(pkg)
}

func getBestPackage() solution {
	packageSizes := []int{5000, 2000, 1000, 500, 250}
	order := 251
	result := optimizePackages(packageSizes, order)
	slices.SortFunc(result, func(a, b solution) int {
		if n := cmp.Compare(a.waste(order), b.waste(order)); n != 0 {
			return n
		}
		return cmp.Compare(a.numberOfPackages(), b.numberOfPackages())
	})

	return result[0]
}

func optimizePackages(packageSizes []int, order int) []solution {
	var result []solution
	currentCombination := solution{make(map[int]int)}
	memo := make(map[int]solution)
	backtrack(packageSizes, order, currentCombination, &result, memo)

	return result
}

func backtrack(packageSizes []int, order int, currentCombination solution, result *[]solution, memo map[int]solution) solution {
	if order <= 0 {
		*result = append(*result, currentCombination)
		return currentCombination
	}
	if _, ok := memo[order]; ok {
		return memo[order]
	}
	for i := 0; i < len(packageSizes); i++ {
		packageSize := packageSizes[i]
		newCombination := solution{make(map[int]int)}
		for k, v := range currentCombination.currentCombination {
			newCombination.currentCombination[k] = v
		}

		newCombination.currentCombination[packageSize] += 1

		memo[order] = backtrack(packageSizes, order-packageSize, newCombination, result, memo)
	}
	return memo[order]
}
