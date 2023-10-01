package domain

type SizedPackage struct {
	Quantity int
	Size     int
}

type CandidatePackages struct {
	CurrentCombination map[int]int
}

func (s CandidatePackages) Waste(order int) int {
	sum := 0
	for key, value := range s.CurrentCombination {
		sum += key * value
	}
	return sum - order
}

func (s CandidatePackages) NumberOfPackages() int {
	sum := 0
	for _, value := range s.CurrentCombination {
		sum += value
	}
	return sum
}
