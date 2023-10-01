package usecase

import (
	"errors"
	"github/ahmedghazey/packaging/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a mock for the PackageService interface
type MockPackageService struct {
	mock.Mock
}

func (m *MockPackageService) GetAllPackages() []*domain.Package {
	args := m.Called()
	return args.Get(0).([]*domain.Package)
}

func (m *MockPackageService) CreatePackage(packages ...*domain.Package) error {
	args := m.Called(packages)
	return args.Error(0)
}

func (m *MockPackageService) GetPackage(id string) (*domain.Package, bool) {
	args := m.Called(id)
	return args.Get(0).(*domain.Package), args.Bool(1)
}

func (m *MockPackageService) UpdatePackage(id string, updatedPackage *domain.Package) (bool, error) {
	args := m.Called(id, updatedPackage)
	return args.Bool(0), args.Error(1)
}

func (m *MockPackageService) DeletePackage(id string) bool {
	args := m.Called(id)
	return args.Bool(0)
}

func TestAddPackages_Execute(t *testing.T) {
	// Create a new instance of the mock PackageService
	mockPackageService := new(MockPackageService)

	// Create an instance of AddPackages with the mock PackageService
	addPackages := NewAddPackages(mockPackageService)

	t.Run("Successful package creation", func(t *testing.T) {
		// Prepare test data
		packages := []*domain.Package{
			&domain.Package{Size: 10},
			&domain.Package{Size: 20},
		}

		// Configure the mock to return no error
		mockPackageService.On("CreatePackage", packages).Return(nil)

		// Call the Execute function
		err := addPackages.Execute(packages)

		// Assert that the error is nil (indicating success)
		assert.NoError(t, err)

		// Verify that the CreatePackage method was called with the correct arguments
		mockPackageService.AssertCalled(t, "CreatePackage", packages)
	})

	t.Run("Failed package creation", func(t *testing.T) {
		// Prepare test data
		packages := []*domain.Package{
			&domain.Package{Size: 30},
		}

		// Configure the mock to return an error
		mockError := errors.New("mock error")
		mockPackageService.On("CreatePackage", packages).Return(mockError)

		// Call the Execute function
		err := addPackages.Execute(packages)

		// Assert that the error is as expected
		assert.EqualError(t, err, "failed to create packages: mock error")

		// Verify that the CreatePackage method was called with the correct arguments
		mockPackageService.AssertCalled(t, "CreatePackage", packages)
	})

	// Cleanup
	mockPackageService.AssertExpectations(t)
}
