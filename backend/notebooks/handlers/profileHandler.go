package handlers

import (
	"encoding/json"
	"net/http"
	"notebooks/notebooks/service"
)

type ProfileHandler struct {
	profileService service.IProfileService
}

func (p *ProfileHandler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	profiles := p.profileService.GetAllProfiles()
	err := json.NewEncoder(w).Encode(profiles)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
