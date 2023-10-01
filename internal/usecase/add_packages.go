package usecase

import (
	"fmt"
	"github/ahmedghazey/packaging/internal/domain"
	"github/ahmedghazey/packaging/internal/service"
)

type AddPackages struct {
	PackagingService service.PackageService
}

func NewAddPackages(packagingService service.PackageService) AddPackages {
	return AddPackages{
		PackagingService: packagingService,
	}
}

func (a AddPackages) Execute(packages []*domain.Package) error {
	err := a.PackagingService.CreatePackage(packages...)
	if err != nil {
		return fmt.Errorf("failed to create packages: %w", err)
	}
	return nil
}
