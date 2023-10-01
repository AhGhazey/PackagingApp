package inmemory

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type Storage struct {
	Items []*Package
	lock  sync.Mutex
}

// NewStorage creates a new instance of Storage.
func NewStorage() *Storage {
	return &Storage{}
}

// Create adds a new Package item to the storage.
func (s *Storage) Create(item *Package) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, existing := range s.Items {
		if existing.Size == item.Size {
			return fmt.Errorf("package with size %d already exists", item.Size)
		}
	}

	s.Items = append(s.Items, item)
	return nil
}

// Get retrieves a Package item from the storage by ID.
func (s *Storage) Get(id uuid.UUID) (*Package, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, item := range s.Items {
		if item.ID == id {
			return item, true
		}
	}

	return nil, false
}

// Update updates a Package item in the storage by ID.
func (s *Storage) Update(id uuid.UUID, updatedPackage *Package) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	for i, item := range s.Items {
		if item.ID == id {
			s.Items[i] = updatedPackage
			return true
		}
	}

	return false
}

// Delete removes a Package item from the storage by ID.
func (s *Storage) Delete(id uuid.UUID) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	for i, item := range s.Items {
		if item.ID == id {
			// Remove the item from the slice by slicing it.
			s.Items = append(s.Items[:i], s.Items[i+1:]...)
			return true
		}
	}

	return false
}

// GetAllPackages fetch all packages
func (s *Storage) GetAllPackages() []*Package {
	s.lock.Lock()
	defer s.lock.Unlock()
	packages := make([]*Package, 0, len(s.Items))
	for _, item := range s.Items {
		packages = append(packages, item)
	}

	return packages
}
