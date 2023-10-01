package repository

import (
	"github.com/google/uuid"
	"github/ahmedghazey/packaging/internal/storage/inmemory"
)

// PackageRepository defines the interface for interacting with the storage.
type PackageRepository interface {
	Create(item *inmemory.Package) error
	Get(id uuid.UUID) (*inmemory.Package, bool)
	Update(id uuid.UUID, updatedPackage *inmemory.Package) bool
	Delete(id uuid.UUID) bool
	GetAllPackages() []*inmemory.Package
}
