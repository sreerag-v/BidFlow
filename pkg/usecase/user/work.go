package userUsecase

import (
	"errors"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/repository/user/interfaces"
	services "github.com/sreerag_v/BidFlow/pkg/usecase/user/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type WorkUsecase struct {
	Repo interfaces.WorkRepo
}

func NewWorkUsecase(repo interfaces.WorkRepo) services.WorkUsecase {
	return &WorkUsecase{
		Repo: repo,
	}
}

func (work *WorkUsecase) ListNewOpening(model domain.ReqWork) error {

	//pass to repository
	err := work.Repo.ListNewOpening(model)
	if err != nil {
		return err
	}

	return nil
}
func (work *WorkUsecase) GetAllListedWorks(id int) ([]models.WorkDetails, error) {

	works, err := work.Repo.GetAllWorksOfAUser(id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	var model []models.WorkDetails

	for _, v := range works {
		details, err := work.Repo.GetDetailsOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find images
		images, err := work.Repo.GetImagesOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		var provider string

		pro_id, err :=work.Repo.FindProviderIdFromWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		if pro_id == 0 {
			provider = "Not Assigned"
		} else {
			provider, err = work.Repo.FindProviderName(pro_id)
			if err != nil {
				return []models.WorkDetails{}, err
			}
		}

		//append
		var result models.WorkDetails
		result.ID = v
		result.Street = details.Street
		result.District = details.District
		result.State = details.State
		result.Profession = details.Profession
		result.User = details.User
		result.Provider = provider
		result.Images = images
		result.WorkStatus = details.WorkStatus

		model = append(model, result)

	}

	return model, nil
}

func (work *WorkUsecase) AddImageOfWork(image string,work_id int)error{
	exist,err:=work.Repo.GetDetailsOfAWork(work_id)
	if err!=nil{
		return err
	}
	if exist.ID == 0 {
		return  errors.New("work does not exist")
	}

	return work.Repo.AddImageOfWork(image,work_id)
}

func (work *WorkUsecase) ListAllCompletedWorks(id int) ([]models.WorkDetails, error) {

	works, err := work.Repo.GetAllCompletedWorksOfAUser(id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	var model []models.WorkDetails

	for _, v := range works {
		details, err := work.Repo.GetDetailsOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find images
		images, err := work.Repo.GetImagesOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		var provider string

		pro_id, err := work.Repo.FindProviderIdFromWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		if pro_id == 0 {
			provider = "Not Assigned"
		} else {
			provider, err = work.Repo.FindProviderName(pro_id)
			if err != nil {
				return []models.WorkDetails{}, err
			}
		}

		//append
		var result models.WorkDetails
		result.ID = v
		result.Street = details.Street
		result.District = details.District
		result.State = details.State
		result.Profession = details.Profession
		result.User = details.User
		result.Provider = provider
		result.Images = images
		result.WorkStatus = details.WorkStatus

		model = append(model, result)

	}

	return model, nil
}

func (work *WorkUsecase) ListAllOngoingWorks(id int) ([]models.WorkDetails, error) {

	works, err := work.Repo.GetAllOngoingWorksOfAUser(id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	var model []models.WorkDetails

	for _, v := range works {
		details, err := work.Repo.GetDetailsOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find images
		images, err := work.Repo.GetImagesOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		var provider string

		pro_id, err := work.Repo.FindProviderIdFromWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		if pro_id == 0 {
			provider = "Not Assigned"
		} else {
			provider, err = work.Repo.FindProviderName(pro_id)
			if err != nil {
				return []models.WorkDetails{}, err
			}
		}

		//append
		var result models.WorkDetails
		result.ID = v
		result.Street = details.Street
		result.District = details.District
		result.State = details.State
		result.Profession = details.Profession
		result.User = details.User
		result.Provider = provider
		result.Images = images
		result.WorkStatus = details.WorkStatus

		model = append(model, result)

	}

	return model, nil
}

func (work *WorkUsecase) WorkDetailsById(id int) (models.WorkDetails, error) {

	details, err := work.Repo.GetDetailsOfAWork(id)
	if err != nil {
		return models.WorkDetails{}, err
	}
	//find image
	images, err := work.Repo.GetImagesOfAWork(id)
	if err != nil {
		return models.WorkDetails{}, err
	}

	var provider string

	pro_id, err := work.Repo.FindProviderIdFromWork(id)
	if err != nil {
		return models.WorkDetails{}, err
	}

	if pro_id == 0 {
		provider = "Not Assigned"
	} else {
		provider, err = work.Repo.FindProviderName(pro_id)
		if err != nil {
			return models.WorkDetails{}, err
		}
	}

	//append
	var result models.WorkDetails
	result.ID = id
	result.Street = details.Street
	result.District = details.District
	result.State = details.State
	result.Profession = details.Profession
	result.User = details.User
	result.Provider = provider
	result.Images = images
	result.WorkStatus = details.WorkStatus

	return result, nil
}