package interfaces

import (
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type ProfileRepo interface{

	AddProfileImage(image string, Uid int)error
	GetImageOfProvider(id int)(string,error)


	CheckIfServiceIsAlreadyRegistered(user_id, service_id int) (bool, error)
	AddService(user_id int, service_id int,profession string) error 
	DeleteService(user_id int, service_id int) error
	CheckIfDistrictIsAlreadyAdded(id, district int) (bool, error)
	AddLocation(id int, district int,dis string) error 
	RemovePreferredLocation(id, district int) error

	GetAllServiceIdsOfAProvider(id int) ([]int, error)
	FindServiceDetailsFromID(id int) (domain.Profession, error)
	GetAllPreferredLocations(id int) ([]int, error)
	GetLocationDetails(id int) (models.GetLocations, error) 
	FindProviderDetails(id int) (domain.Provider, error)
	GetRatingsOfAllRecordsOfAProvider(id int) ([]int, error) 


	GetServiceFromSelected(int)(domain.ProProfileService,error)
	GetLocationFromSelected(int)(domain.ProProfileLocation,error)
}