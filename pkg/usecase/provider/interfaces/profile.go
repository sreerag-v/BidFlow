package interfaces

import "github.com/sreerag_v/BidFlow/pkg/utils/models"

type ProfileUsecase interface{
	AddProfileImage(string,int)error
	AddService(int, int) error
	DeleteService(user_id int, service_id int) error
	AddPreferredWorkingLocation(id int, district int) error 
	RemovePreferredLocation(id, district int) error

	GetMyServices(id int) ([]models.GetServices, error)
	GetAllPreferredLocations(id int) ([]models.GetLocations, error)
	GetDetailsOfProviders(id int) (models.ProviderProfile, error)

	GetProDetails(id int)(models.ProviderProfile, error)
}