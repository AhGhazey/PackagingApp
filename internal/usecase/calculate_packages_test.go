package usecase

import (
	"github.com/stretchr/testify/assert"
	"github/ahmedghazey/packaging/internal/domain"
	"testing"
)

func TestCalculatePackages_Execute(t *testing.T) {
	// Create a new instance of the mock PackagingService
	mockPackagingService := new(MockPackageService)

	// Create an instance of CalculatePackages with the mock PackagingService
	calculatePackages := CalculatePackages{PackagingService: mockPackagingService}

	t.Run("Successful package calculation", func(t *testing.T) {
		// Prepare test data
		numberOfItems := 10

		// Define mock data for GetAllPackages
		mockPackages := []*domain.Package{
			{Size: 5, Id: "00000000-0000-0000-0000-000000000001"},
			{Size: 3, Id: "00000000-0000-0000-0000-000000000002"},
			{Size: 2, Id: "00000000-0000-0000-0000-000000000003"},
		}

		// Configure the mock to return the mockPackages data
		mockPackagingService.On("GetAllPackages").Return(mockPackages)

		// Call the Execute function
		result := calculatePackages.Execute(numberOfItems)

		// Verify that the result contains the expected packages
		expectedResult := []*domain.SizedPackage{
			{Size: 5, Quantity: 2},
		}
		assert.Equal(t, expectedResult, result)

		// Verify that the GetAllPackages method was called
		mockPackagingService.AssertCalled(t, "GetAllPackages")
	})

	// Add more test cases as needed to cover different scenarios

	// Cleanup
	mockPackagingService.AssertExpectations(t)
}
