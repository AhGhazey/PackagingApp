package rest

import (
	"encoding/json"
	"github/ahmedghazey/packaging/internal/domain"
	"github/ahmedghazey/packaging/internal/service"
	"github/ahmedghazey/packaging/internal/usecase"
	"net/http"
)

type Package struct {
	Size int `json:"size"`
}
type AddPackagesRequest struct {
	Packages []Package `json:"packages"`
}
type AddPackagesResponse struct {
	Message string `json:"message"`
}

// AddPackages
// @Summary Add packages
// @Description Add packages to the system
// @Tags Packages
// @Accept json
// @Produce json
// @Param request body AddPackagesRequest true "Request body with packages to add"
// @Success 200 {object} AddPackagesResponse "Packages added successfully"
// @Failure 400 {object} string "Invalid request format or package size"
// @Failure 500 {object} string "Internal server error"
// @Router /add-packages [post]
func AddPackages(packagingService service.PackageService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var addPackageRequest AddPackagesRequest
		err := json.NewDecoder(r.Body).Decode(&addPackageRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		addPackagesUsecase := usecase.NewAddPackages(packagingService)
		packages := make([]*domain.Package, 0, len(addPackageRequest.Packages))
		for _, pkg := range addPackageRequest.Packages {
			if pkg.Size <= 0 {
				http.Error(w, "Package size must be a positive integer greater than 0", http.StatusBadRequest)
				return
			}
			packages = append(packages, &domain.Package{Size: pkg.Size})
		}
		err = addPackagesUsecase.Execute(packages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := AddPackagesResponse{
			Message: "Packages added",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusOK)
	}
}
