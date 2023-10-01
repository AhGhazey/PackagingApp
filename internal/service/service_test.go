package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github/ahmedghazey/packaging/internal/domain"
	"github/ahmedghazey/packaging/internal/storage/inmemory"
	"testing"
)

// MockRepository is a mock implementation of PackageRepository
type MockPackageRepository struct {
	mock.Mock
}

func (m *MockPackageRepository) Create(item *inmemory.Package) error {
	args := m.Called(item)
	return args.Error(0)
}
func (m *MockPackageRepository) Get(id uuid.UUID) (*inmemory.Package, bool) {
	args := m.Called(id)
	return args.Get(0).(*inmemory.Package), args.Bool(1)
}
func (m *MockPackageRepository) Update(id uuid.UUID, updatedPackage *inmemory.Package) bool {
	args := m.Called(id, updatedPackage)
	return args.Bool(0)
}
func (m *MockPackageRepository) Delete(id uuid.UUID) bool {
	args := m.Called(id)
	return args.Bool(0)
}
func (m *MockPackageRepository) GetAllPackages() []*inmemory.Package {
	args := m.Called()
	return args.Get(0).([]*inmemory.Package)
}

func TestDomainToStorage(t *testing.T) {
	// Valid UUID string
	validUUIDStr := "00000000-0000-0000-0000-000000000000"
	validUUID, _ := uuid.Parse(validUUIDStr)

	testCases := []struct {
		name           string
		domainPackage  *domain.Package
		expectedResult *inmemory.Package
		expectedError  error
	}{
		{
			name: "Valid Input",
			domainPackage: &domain.Package{
				Id:   validUUIDStr,
				Size: 10,
			},
			expectedResult: &inmemory.Package{
				ID:   validUUID,
				Size: 10,
			},
			expectedError: nil,
		},
		{
			name: "Invalid UUID Format",
			domainPackage: &domain.Package{
				Id:   "invalid-uuid-format",
				Size: 20,
			},
			expectedResult: nil,
			expectedError:  fmt.Errorf("invalid UUID"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := domainToStorage(testCase.domainPackage)
			assert.Equal(t, testCase.expectedResult, result, "Result does not match expected")
			assert.Equal(t, testCase.expectedError, err, "Error does not match expected")
		})
	}
}

func TestService_CreatePackage(t *testing.T) {
	testCases := []struct {
		name           string
		pkg            *domain.Package
		expectedError  error
		expectedCalled bool
	}{
		{
			name: "Successful package creation",
			pkg: &domain.Package{
				Id:   "00000000-0000-0000-0000-000000000001",
				Size: 10,
			},
			expectedError:  nil,
			expectedCalled: true,
		},
		{
			name: "Failed package creation",
			pkg: &domain.Package{
				Id:   "00000000-0000-0000-0000-000000000002",
				Size: 20,
			},
			expectedError:  fmt.Errorf("failed to create package: %w", fmt.Errorf("mock repository error")),
			expectedCalled: true,
		},
		{
			name: "Failed package creation - Invalid UUID format",
			pkg: &domain.Package{
				Id:   "invalid-uuid-format",
				Size: 20,
			},
			expectedError:  fmt.Errorf("failed to create package: %w", fmt.Errorf("mock repository error")),
			expectedCalled: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := new(MockPackageRepository)
			service := NewService(mockRepository)
			expectedStoragePackage, _ := domainToStorage(tc.pkg)
			mockRepository.On("Create", expectedStoragePackage).Return(tc.expectedError)

			err := service.CreatePackage(tc.pkg)
			if tc.expectedError != nil {
				assert.Equal(t, fmt.Errorf("failed to create package: %w", tc.expectedError), err)
			} else {
				assert.Nil(t, err)
			}

			mockRepository.AssertCalled(t, "Create", expectedStoragePackage)

			if tc.expectedCalled {
				mockRepository.AssertExpectations(t)
			} else {
				mockRepository.AssertNotCalled(t, "Create")
			}
		})
	}
}

func TestService_GetPackage(t *testing.T) {
	testCases := []struct {
		name           string
		id             string
		expectedResult *domain.Package
		expectedFound  bool
		expectedCalled bool
	}{
		{
			name: "Valid package found",
			id:   "00000000-0000-0000-0000-000000000001",
			expectedResult: &domain.Package{
				Id:   "00000000-0000-0000-0000-000000000001",
				Size: 10,
			},
			expectedFound:  true,
			expectedCalled: true,
		},
		{
			name:           "Invalid UUID format",
			id:             "invalid-uuid-format",
			expectedResult: nil,
			expectedFound:  false,
			expectedCalled: false,
		},
		{
			name:           "Package not found",
			id:             "00000000-0000-0000-0000-000000000002",
			expectedResult: nil,
			expectedFound:  false,
			expectedCalled: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, _ := uuid.Parse(tc.id)

			mockRepository := new(MockPackageRepository)

			service := NewService(mockRepository)

			var expectedResult *inmemory.Package
			if tc.expectedResult != nil {
				expectedResult, _ = domainToStorage(tc.expectedResult)
			}

			mockRepository.On("Get", id).Return(expectedResult, tc.expectedFound)

			result, found := service.GetPackage(tc.id)

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedFound, found)

			if tc.expectedCalled {
				mockRepository.AssertExpectations(t)
			} else {
				mockRepository.AssertNotCalled(t, "Get")
			}

		})
	}
}

func TestService_UpdatePackage(t *testing.T) {
	// Create a new instance of the mock PackageRepository
	mockRepository := new(MockPackageRepository)

	// Create an instance of the Service with the mock PackageRepository
	service := NewService(mockRepository)

	testCases := []struct {
		name               string
		id                 string
		updatedPackage     *domain.Package
		mockReturn         bool
		expectedResult     bool
		expectedError      error
		expectedCalledOnce bool
	}{
		{
			name: "Invalid UUID format",
			id:   "invalid-uuid-format",
			updatedPackage: &domain.Package{
				Id:   "00000000-0000-0000-0000-000000000002",
				Size: 20,
			},
			mockReturn:         false,
			expectedResult:     false,
			expectedError:      errors.New("invalid UUID"),
			expectedCalledOnce: false,
		},
		{
			name: "Valid package update",
			id:   "00000000-0000-0000-0000-000000000001",
			updatedPackage: &domain.Package{
				Id:   "00000000-0000-0000-0000-000000000001",
				Size: 10,
			},
			mockReturn:         true,
			expectedResult:     true,
			expectedError:      nil,
			expectedCalledOnce: true,
		},
		{
			name: "Package update failed",
			id:   "00000000-0000-0000-0000-000000000004",
			updatedPackage: &domain.Package{
				Id:   "00000000-0000-0000-0000-000000000004",
				Size: 20,
			},
			mockReturn:         false,
			expectedResult:     false,
			expectedError:      nil,
			expectedCalledOnce: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Prepare test data
			id, _ := uuid.Parse(tc.id)
			pkg, _ := domainToStorage(tc.updatedPackage)
			mockRepository.On("Update", id, pkg).Return(tc.mockReturn)

			// Call the UpdatePackage function
			result, err := service.UpdatePackage(tc.id, tc.updatedPackage)

			// Assert the result and error
			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedError, err)

			// Verify that the Update method of the mock repository was called
			if tc.expectedCalledOnce {
				mockRepository.AssertCalled(t, "Update", id, mock.AnythingOfType("*inmemory.Package"))
			} else {
				mockRepository.AssertNotCalled(t, "Update", 0)
			}
		})
	}
}

func TestService_GetAllPackages(t *testing.T) {
	// Define a function to create a new instance of the mock repository
	newMockRepository := func() *MockPackageRepository {
		return new(MockPackageRepository)
	}

	firstId, _ := uuid.Parse("00000000-0000-0000-0000-000000000001")
	secondId, _ := uuid.Parse("00000000-0000-0000-0000-000000000002")
	thirdId, _ := uuid.Parse("00000000-0000-0000-0000-000000000003")

	testCases := []struct {
		name                 string
		mockReturn           []*inmemory.Package
		expectedResultLength int
		expectedFirstID      *uuid.UUID
		expectedSecondID     *uuid.UUID
	}{
		{
			name:                 "Empty package retrieval",
			mockReturn:           []*inmemory.Package{},
			expectedResultLength: 0,
			expectedFirstID:      nil,
			expectedSecondID:     nil,
		},
		{
			name: "Valid package retrieval",
			mockReturn: []*inmemory.Package{
				{ID: firstId, Size: 10},
				{ID: secondId, Size: 20},
				{ID: thirdId, Size: 15},
			},
			expectedResultLength: 3,
			expectedFirstID:      &secondId,
			expectedSecondID:     &thirdId,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := newMockRepository()
			service := NewService(mockRepository)

			mockRepository.On("GetAllPackages").Return(tc.mockReturn)

			result := service.GetAllPackages()

			assert.Len(t, result, tc.expectedResultLength)
			if len(result) > 0 {
				assert.Equal(t, tc.expectedFirstID.String(), result[0].Id)
			}
			if len(result) > 1 {
				assert.Equal(t, tc.expectedSecondID.String(), result[1].Id)
			}

			mockRepository.AssertCalled(t, "GetAllPackages")
		})
	}
}

func TestService_DeletePackage(t *testing.T) {
	testCases := []struct {
		name               string
		id                 string
		mockReturn         bool
		expectedResult     bool
		expectedCalledOnce bool
	}{
		{
			name:               "Valid package deletion",
			id:                 "00000000-0000-0000-0000-000000000001",
			mockReturn:         true,
			expectedResult:     true,
			expectedCalledOnce: true,
		},
		{
			name:               "Invalid UUID format",
			id:                 "invalid-uuid-format",
			mockReturn:         false,
			expectedResult:     false,
			expectedCalledOnce: false,
		},
		{
			name:               "Failed package deletion",
			id:                 "00000000-0000-0000-0000-000000000002",
			mockReturn:         false,
			expectedResult:     false,
			expectedCalledOnce: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepository := new(MockPackageRepository)
			service := NewService(mockRepository)

			expectedUUID, _ := uuid.Parse(tc.id)
			mockRepository.On("Delete", expectedUUID).Return(tc.mockReturn)

			result := service.DeletePackage(tc.id)

			assert.Equal(t, tc.expectedResult, result)

			if tc.expectedCalledOnce {
				mockRepository.AssertCalled(t, "Delete", expectedUUID)
			} else {
				mockRepository.AssertNotCalled(t, "Delete")
			}
		})
	}
}
