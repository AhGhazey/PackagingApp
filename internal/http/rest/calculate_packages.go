package rest

import (
	"encoding/json"
	"github/ahmedghazey/packaging/internal/service"
	"github/ahmedghazey/packaging/internal/usecase"
	"net/http"
)

type CalculatePackagesRequest struct {
	Amount int `json:"amount"`
}
type SizedPackage struct {
	Quantity int `json:"quantity"`
	Size     int `json:"size"`
}
type CalculatePackagesResponse struct {
	Packages []*SizedPackage `json:"packages"`
}

// CalculatePackages
// @Summary Calculate required packages
// @Description Calculate the minimum number of packages required for a given amount of items
// @Tags Packages
// @Accept json
// @Produce json
// @Param request body CalculatePackagesRequest true "Request body with the amount of items"
// @Success 200 {object} CalculatePackagesResponse "Minimum number of packages calculated successfully"
// @Failure 400 {object} string "Invalid request format or amount"
// @Failure 500 {object} string "Internal server error"
// @Router /calculate-packages [post]
func CalculatePackages(packagingService service.PackageService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var calculatePackagesRequest CalculatePackagesRequest
		err := json.NewDecoder(r.Body).Decode(&calculatePackagesRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if calculatePackagesRequest.Amount <= 0 {
			http.Error(w, "Amount must be a positive integer greater than 0", http.StatusBadRequest)
			return
		}

		calculatePackagesUsecase := usecase.NewCalculatePackages(packagingService)
		sizedPackages := calculatePackagesUsecase.Execute(calculatePackagesRequest.Amount)
		response := CalculatePackagesResponse{}
		for _, sizedPackage := range sizedPackages {
			response.Packages = append(response.Packages, &SizedPackage{
				Quantity: sizedPackage.Quantity,
				Size:     sizedPackage.Size,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusOK)
	}
}
