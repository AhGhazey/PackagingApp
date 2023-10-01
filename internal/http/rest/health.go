package rest

import (
	"github.com/go-chi/render"
	"github/ahmedghazey/packaging/pkg/logging"
	"net/http"
)

// Health
// @Health
// @Summary get request to check service health
// @Description  get request to check service health
// @Produce      json
// @Success      200      {object}  HealthResponse
// @Failure      500      {string}  string         "fail"
// @Failure      408      {string}  string         "fail"
// @Router       /health [get]
func Health() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logging.Logger.WithContext(r.Context()).Info("Received request to check service health")
		res := NewHealthResponse(true, "alive")
		if err := render.Render(w, r, res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type HealthResponse struct {
	IsAlive bool   `json:"isAlive"`
	Message string `json:"msg"`
}

func (*HealthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewHealthResponse(flag bool, msg string) *HealthResponse {
	return &HealthResponse{
		IsAlive: flag,
		Message: msg,
	}
}
