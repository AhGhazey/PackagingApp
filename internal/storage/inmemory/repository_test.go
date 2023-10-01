package inmemory

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {

	storage := &Storage{}

	tests := []struct {
		name     string
		input    *Package
		expected []*Package
	}{
		{
			name:     "Add single package",
			input:    &Package{ID: uuid.New(), Size: 10},
			expected: []*Package{{ID: uuid.New(), Size: 10}},
		},
		{
			name:     "Add multiple packages",
			input:    &Package{ID: uuid.New(), Size: 20},
			expected: []*Package{{ID: uuid.New(), Size: 10}, {ID: uuid.New(), Size: 20}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			storage.Create(test.input)

			assert.Len(t, storage.Items, len(test.expected), "Storage should have the expected number of items")
		})
	}
}
func TestGet(t *testing.T) {

	s := &Storage{}

	p1 := &Package{ID: uuid.New(), Size: 10}
	p2 := &Package{ID: uuid.New(), Size: 20}
	s.Items = append(s.Items, p1, p2)

	tests := []struct {
		name         string
		searchID     uuid.UUID
		expectedPkg  *Package
		expectedBool bool
	}{
		{
			name:         "Item found",
			searchID:     p1.ID,
			expectedPkg:  p1,
			expectedBool: true,
		},
		{
			name:         "Item not found",
			searchID:     uuid.New(),
			expectedPkg:  nil,
			expectedBool: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, found := s.Get(test.searchID)

			assert.Equal(t, test.expectedPkg, result, "Returned package does not match expected")
			assert.Equal(t, test.expectedBool, found, "Found flag does not match expected")
		})
	}
}
func TestUpdate(t *testing.T) {
	s := &Storage{}

	p1 := &Package{ID: uuid.New(), Size: 10}
	p2 := &Package{ID: uuid.New(), Size: 20}
	s.Items = append(s.Items, p1, p2)

	tests := []struct {
		name            string
		updateID        uuid.UUID
		updatedPackage  *Package
		expectedSuccess bool
		expectedItems   []*Package
	}{
		{
			name:            "Update existing item",
			updateID:        p1.ID,
			updatedPackage:  &Package{ID: p1.ID, Size: 30},
			expectedSuccess: true,
			expectedItems:   []*Package{{ID: p1.ID, Size: 30}, {ID: p2.ID, Size: 20}},
		},
		{
			name:            "Update non-existing item",
			updateID:        uuid.New(),
			updatedPackage:  &Package{ID: uuid.New(), Size: 40},
			expectedSuccess: false,
			expectedItems:   []*Package{{ID: p1.ID, Size: 30}, {ID: p2.ID, Size: 20}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			success := s.Update(test.updateID, test.updatedPackage)

			assert.Equal(t, test.expectedSuccess, success, "Update success status does not match expected")
			assert.Equal(t, test.expectedItems, s.Items, "Storage items do not match expected after update")

		})
	}
}
func TestDelete(t *testing.T) {
	s := &Storage{}

	p1 := &Package{ID: uuid.New(), Size: 10}
	p2 := &Package{ID: uuid.New(), Size: 20}
	s.Items = append(s.Items, p1, p2)

	tests := []struct {
		name            string
		deleteID        uuid.UUID
		expectedSuccess bool
		expectedItems   []*Package
	}{
		{
			name:            "Delete existing item",
			deleteID:        p1.ID,
			expectedSuccess: true,
			expectedItems:   []*Package{{ID: p2.ID, Size: 20}},
		},
		{
			name:            "Delete non-existing item",
			deleteID:        uuid.New(),
			expectedSuccess: false,
			expectedItems:   []*Package{{ID: p2.ID, Size: 20}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			success := s.Delete(test.deleteID)

			assert.Equal(t, test.expectedSuccess, success, "Delete success status does not match expected")
			assert.Equal(t, test.expectedItems, s.Items, "Storage items do not match expected after delete")

		})
	}
}
func TestGetAllPackages(t *testing.T) {
	s := &Storage{}

	p1 := &Package{ID: uuid.New(), Size: 10}
	p2 := &Package{ID: uuid.New(), Size: 20}
	s.Items = append(s.Items, p1, p2)

	packages := s.GetAllPackages()

	assert.Len(t, packages, len(s.Items), "Number of packages returned by GetAllPackages does not match the storage")

	for i, item := range packages {
		assert.Equal(t, s.Items[i], item, "Package at index %d does not match the storage", i)
	}
}
