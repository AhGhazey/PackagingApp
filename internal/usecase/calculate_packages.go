package usecase

import (
	"cmp"
	"github/ahmedghazey/packaging/internal/domain"
	"github/ahmedghazey/packaging/internal/service"
	"slices"
)

type CalculatePackages struct {
	PackagingService service.PackageService
}

func NewCalculatePackages(packagingService service.PackageService) CalculatePackages {
	return CalculatePackages{
		PackagingService: packagingService,
	}
}

func (c CalculatePackages) Execute(numberOfItems int) []*domain.SizedPackage {
	existingPackages := c.PackagingService.GetAllPackages() //return data sorted descending
	packages := make([]*domain.SizedPackage, 0, len(existingPackages))
	result := optimizePackages(existingPackages, numberOfItems)
	slices.SortFunc(result, func(a, b domain.CandidatePackages) int {
		if n := cmp.Compare(a.Waste(numberOfItems), b.Waste(numberOfItems)); n != 0 {
			return n
		}
		return cmp.Compare(a.NumberOfPackages(), b.NumberOfPackages())
	})

	for k, v := range result[0].CurrentCombination {
		packages = append(packages, &domain.SizedPackage{Size: k, Quantity: v})
	}
	return packages
}

func optimizePackages(packages []*domain.Package, order int) []domain.CandidatePackages {
	var result []domain.CandidatePackages
	currentCombination := domain.CandidatePackages{make(map[int]int)}
	memo := make(map[int]domain.CandidatePackages)
	backtrack(packages, order, currentCombination, &result, memo)

	return result
}

func backtrack(packageSizes []*domain.Package, order int,
	currentCombination domain.CandidatePackages,
	result *[]domain.CandidatePackages,
	memo map[int]domain.CandidatePackages) domain.CandidatePackages {

	if order <= 0 {
		*result = append(*result, currentCombination)
		return currentCombination
	}
	if _, ok := memo[order]; ok {
		return memo[order]
	}
	for i := 0; i < len(packageSizes); i++ {
		packageSize := packageSizes[i].Size
		newCombination := domain.CandidatePackages{make(map[int]int)}
		for k, v := range currentCombination.CurrentCombination {
			newCombination.CurrentCombination[k] = v
		}

		newCombination.CurrentCombination[packageSize] += 1

		memo[order] = backtrack(packageSizes, order-packageSize, newCombination, result, memo)
	}
	return memo[order]
}
