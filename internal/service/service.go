package service

import (
	"cmp"
	"fmt"
	"github.com/google/uuid"
	"github/ahmedghazey/packaging/internal/domain"
	"github/ahmedghazey/packaging/internal/repository"
	"github/ahmedghazey/packaging/internal/storage/inmemory"
	"slices"
)

type PackageService interface {
	CreatePackage(...*domain.Package) error
	GetPackage(id string) (*domain.Package, bool)
	UpdatePackage(id string, updatedPackage *domain.Package) (bool, error)
	DeletePackage(id string) bool
	GetAllPackages() []*domain.Package
}

var _ PackageService = (*Service)(nil)

type Service struct {
	repository repository.PackageRepository
}

func NewService(repository repository.PackageRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreatePackage(items ...*domain.Package) error {
	for _, pkg := range items {
		if pkg.Id == "" {
			pkg.Id = uuid.New().String()
		}

		item, _ := domainToStorage(pkg)
		err := s.repository.Create(item)
		if err != nil {
			return fmt.Errorf("failed to create package: %w", err)
		}
	}
	return nil
}
func (s *Service) GetPackage(id string) (*domain.Package, bool) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, false
	}

	storagePkg, found := s.repository.Get(uuidID)
	if !found {
		return nil, false
	}

	return storageToDomain(storagePkg), true
}
func (s *Service) UpdatePackage(id string, updatedPackage *domain.Package) (bool, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return false, fmt.Errorf("invalid UUID")
	}
	storageItem, err := domainToStorage(updatedPackage)
	if err != nil {
		return false, fmt.Errorf("failed to convert package to storage")
	}
	return s.repository.Update(uuidID, storageItem), nil
}
func (s *Service) DeletePackage(id string) bool {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return false
	}

	return s.repository.Delete(uuidID)
}
func (s *Service) GetAllPackages() []*domain.Package {
	storagePackages := s.repository.GetAllPackages()
	domainPackages := make([]*domain.Package, 0, len(storagePackages))

	for _, storagePkg := range storagePackages {
		domainPackages = append(domainPackages, storageToDomain(storagePkg))
	}
	slices.SortFunc(domainPackages, func(a, b *domain.Package) int {
		return cmp.Compare(b.Size, a.Size)
	})
	return domainPackages
}

func domainToStorage(domainPkg *domain.Package) (*inmemory.Package, error) {
	uuidID, err := uuid.Parse(domainPkg.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID")
	}
	return &inmemory.Package{
		ID:   uuidID,
		Size: domainPkg.Size,
	}, nil
}
func storageToDomain(storagePkg *inmemory.Package) *domain.Package {
	return &domain.Package{
		Id:   storagePkg.ID.String(),
		Size: storagePkg.Size,
	}
}
