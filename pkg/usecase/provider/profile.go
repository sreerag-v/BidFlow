package providerUsecase

import (
	"errors"

	"github.com/sreerag_v/BidFlow/pkg/repository/provider/interfaces"
	service "github.com/sreerag_v/BidFlow/pkg/usecase/provider/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type ProfileUsecase struct{
	Repo interfaces.ProfileRepo
}

func NewProfileUsecase(repo interfaces.ProfileRepo)service.ProfileUsecase{
	return &ProfileUsecase{
		Repo: repo,
	}
}

func (pr *ProfileUsecase) AddService(user_id int, service_id int) error {
	//check if the service already exists if not add the service
	exists, err := pr.Repo.CheckIfServiceIsAlreadyRegistered(user_id, service_id)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("already exists")
	}

	if err := pr.Repo.AddService(user_id, service_id); err != nil {
		return err
	}

	return nil
}

func (pr *ProfileUsecase) DeleteService(user_id int, service_id int) error {
	//check if the service already exists is yes delete the service
	exists, err := pr.Repo.CheckIfServiceIsAlreadyRegistered(user_id, service_id)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("service not found")
	}

	if err := pr.Repo.DeleteService(user_id, service_id); err != nil {
		return err
	}

	return nil
}

func (pr*ProfileUsecase) AddPreferredWorkingLocation(id int, district int) error {
	exists, err := pr.Repo.CheckIfDistrictIsAlreadyAdded(id, district)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("already exists")
	}

	if err := pr.Repo.AddLocation(id, district); err != nil {
		return err
	}

	return nil
}

func (pr *ProfileUsecase) RemovePreferredLocation(id int, district int) error {

	if err := pr.Repo.RemovePreferredLocation(id, district); err != nil {
		return err
	}

	return nil
}

func (pr *ProfileUsecase) GetMyServices(id int) ([]models.GetServices, error) {
	var model []models.GetServices
	//find serviceIDs
	all_id, err := pr.Repo.GetAllServiceIdsOfAProvider(id)
	if err != nil {
		return []models.GetServices{}, err
	}
	//find category
	for _, v := range all_id {
		service := models.GetServices{}
		service.ID = v
		details, err := pr.Repo.FindServiceDetailsFromID(v)
		if err != nil {
			return []models.GetServices{}, err
		}
		service.ServiceName = details.Profession

		service.Category_id = details.CategoryID
		model = append(model, service)
	}

	//find service name
	return model, nil
}

func (pr ProfileUsecase) GetAllPreferredLocations(id int) ([]models.GetLocations, error) {
	var model []models.GetLocations
	//find serviceIDs
	locations, err := pr.Repo.GetAllPreferredLocations(id)
	if err != nil {
		return []models.GetLocations{}, err
	}
	//find category
	for _, v := range locations {
		details, err := pr.Repo.GetLocationDetails(v)
		if err != nil {
			return []models.GetLocations{}, err
		}
		details.ID = v
		model = append(model, details)
	}

	//find service name
	return model, nil
}

func (pr *ProfileUsecase) GetDetailsOfProviders(id int) (models.ProviderDetailsForUser, error) {

	var model models.ProviderDetailsForUser
	//find details of providers
	details, err := pr.Repo.FindProviderDetails(id)
	if err != nil {
		return models.ProviderDetailsForUser{}, err
	}

	model.ID = id
	model.Name = details.Name
	model.Email = details.Email
	model.Phone = details.Phone

	ratings, err := pr.Repo.GetRatingsOfAllRecordsOfAProvider(id)
	if err != nil {
		return models.ProviderDetailsForUser{}, err
	}

	var sum int
	//find average rating of provider
	for _, v := range ratings {
		sum = sum + v
	}

	length := len(ratings)
	if length == 0 {
		length = 1
	}

	average := sum / length

	model.AverageRating = average

	return model, nil
}