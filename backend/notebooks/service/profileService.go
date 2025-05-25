package service

import (
	"notebooks/notebooks/domain"
	"notebooks/notebooks/dto"
)

type IProfileService interface {
	GetAllProfiles() []dto.Profile
}
type ProfileService struct {
	repo domain.INotebookRepo
}

/******* Profile *******/
func (p *ProfileService) GetAllProfiles() []dto.Profile {
	return p.repo.GetProfilesByUsername("Vini")
}
