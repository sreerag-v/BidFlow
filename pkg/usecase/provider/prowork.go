package providerUsecase

import (
	"errors"
	"fmt"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	sent "github.com/sreerag_v/BidFlow/pkg/notification/sender"
	"github.com/sreerag_v/BidFlow/pkg/repository/provider/interfaces"
	services "github.com/sreerag_v/BidFlow/pkg/usecase/provider/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
)

type ProWorkUsecase struct {
	Repo        interfaces.ProWorkRepo
	ProfileRepo interfaces.ProfileRepo
}

func NewProWorkUsecase(repo interfaces.ProWorkRepo, profile interfaces.ProfileRepo) services.ProWorkUsecase {
	return &ProWorkUsecase{
		Repo:        repo,
		ProfileRepo: profile,
	}
}

func (w *ProWorkUsecase) GetAllLeads(pro_id int, page models.PageNation) ([]models.WorkDetails, error) {

	//get providers preferred services
	services, err := w.ProfileRepo.GetAllServiceIdsOfAProvider(pro_id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	fmt.Println("services", services)

	//get providers preffered locations
	locations, err := w.ProfileRepo.GetAllPreferredLocations(pro_id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	fmt.Println("locations", locations)

	//get work ids by these preferences
	var works []int
	for _, service := range services {
		for _, location := range locations {
			lead, err := w.Repo.GetLeadByServiceAndLocation(service, location)
			if err != nil {
				return []models.WorkDetails{}, err
			}
			fmt.Println("lead", lead)
			works = append(works, lead...)

		}
	}

	var model []models.WorkDetails
	for _, v := range works {
		//find details
		details, err := w.Repo.GetDetailsOfAWork(v, page)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		if details.ID == 0 {
			return nil, nil
		}

		fmt.Println("details:", details)
		//find images
		images, err := w.Repo.GetImagesOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		exists, err := w.Repo.CheckIfAlreadyBidded(v, pro_id)
		if err != nil {
			return []models.WorkDetails{}, err
		}

		//append
		var result models.WorkDetails
		result.ID = v
		result.Street = details.Street
		result.District = details.District
		result.State = details.State
		result.Profession = details.Profession
		result.User = details.User
		result.Images = images
		result.Participation = exists
		result.WorkStatus = details.WorkStatus
		model = append(model, result)
	}

	//pack into model and return the model

	fmt.Println("model", model)
	return model, nil
}

func (w *ProWorkUsecase) ViewLeads(work_id int, page models.PageNation) (models.WorkDetails, error) {
	details, err := w.Repo.GetDetailsOfAWork(work_id, page)
	if err != nil {
		return models.WorkDetails{}, err
	}
	//find images
	images, err := w.Repo.GetImagesOfAWork(work_id)
	if err != nil {
		return models.WorkDetails{}, err
	}
	//append
	var result models.WorkDetails
	result.ID = work_id
	result.Street = details.Street
	result.District = details.District
	result.State = details.State
	result.Profession = details.Profession
	result.User = details.User
	result.Images = images
	result.WorkStatus = details.WorkStatus

	return result, nil
}

func (w *ProWorkUsecase) PlaceBid(model models.PlaceBid) error {
	exist, err := w.Repo.FindWorkExistOrNot(model.WorkID)
	if err != nil {
		return err
	}

	if exist.ID == 0 {
		return errors.New("Work Not Found For Bidding")
	}
	checkbid, err := w.Repo.CheckBidExistOrNot(model.WorkID)
	if err != nil {
		return err
	}

	if checkbid.ID != 0 {
		return errors.New("Bid Already Palced On This Work")
	}

	if checkbid.IsDeleted {
		return errors.New("Bid Already Enrolled")
	}

	err = w.Repo.PlaceBid(model, exist.UserID)
	if err != nil {
		return err
	} else {
		user, err := w.Repo.FindUserByUid(exist.UserID)
		if err != nil {
			return err
		}
		// pass those into the Sender()
		sent.Sent(user.Email, user.Name)
	}

	return nil
}

func (w *ProWorkUsecase) ReplaceBidWithNewBid(model models.PlaceBid) error {
	exist, err := w.Repo.FindWorkExistOrNot(model.WorkID)
	if err != nil {
		return err
	}

	if exist.ID == 0 {
		return errors.New("Work Not Found For Replace Bid")
	}
	checkbid, err := w.Repo.CheckBidExistOrNot(model.WorkID)
	if err != nil {
		return err
	}
	if checkbid.IsDeleted {
		return errors.New("Bid Already Enrolled")
	}

	err = w.Repo.ReplaceBidWithNewBid(model, exist.UserID)
	if err != nil {
		return err
	} else {
		user, err := w.Repo.FindUserByUid(exist.UserID)
		if err != nil {
			return err
		}
		// pass those into the Sender()
		sent.Sent(user.Email, user.Name)
	}

	return nil
}

func (w *ProWorkUsecase) GetAllOtherBidsOnTheLeads(work_id int) ([]models.BidDetails, error) {
	bids, err := w.Repo.GetAllOtherBidsOnTheLeads(work_id)
	if err != nil {
		return []models.BidDetails{}, err
	}

	return bids, nil
}

func (w *ProWorkUsecase) GetWorksOfAProvider(pro_id int, page models.PageNation) ([]models.WorkDetails, error) {

	provider, err := w.Repo.FindProviderName(pro_id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	works, err := w.Repo.GetAllWorksOfAProvider(pro_id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	var model []models.WorkDetails

	for _, v := range works {
		details, err := w.Repo.GetDetailsOfAWork(v, page)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find images
		images, err := w.Repo.GetImagesOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
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

func (w *ProWorkUsecase) GetAllOnGoingWorks(pro_id int, page models.PageNation) ([]models.WorkDetails, error) {

	provider, err := w.Repo.FindProviderName(pro_id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	works, err := w.Repo.GetCommittedWorksOfAProvider(pro_id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	var model []models.WorkDetails

	for _, v := range works {
		details, err := w.Repo.GetDetailsOfAWork(v, page)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find images
		images, err := w.Repo.GetImagesOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
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

func (w *ProWorkUsecase) GetCompletedWorks(pro_id int, page models.PageNation) ([]models.WorkDetails, error) {

	provider, err := w.Repo.FindProviderName(pro_id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	works, err := w.Repo.GetCompletedWorksOfAProvider(pro_id)
	if err != nil {
		return []models.WorkDetails{}, err
	}

	var model []models.WorkDetails

	for _, v := range works {
		details, err := w.Repo.GetDetailsOfAWork(v, page)
		if err != nil {
			return []models.WorkDetails{}, err
		}
		//find images
		images, err := w.Repo.GetImagesOfAWork(v)
		if err != nil {
			return []models.WorkDetails{}, err
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

func (w *ProWorkUsecase) GetAllAcceptedBids(pro_id int, page models.PageNation) ([]domain.AcceptedBidRes, error) {
	exist, err := w.ProfileRepo.FindProviderDetails(pro_id)

	if err != nil {
		return []domain.AcceptedBidRes{}, err
	}

	if exist.ID == 0 {
		return []domain.AcceptedBidRes{}, errors.New("Provider Does Not Exists")
	}

	bids, err := w.Repo.GetAllAcceptedBids(pro_id, page)
	if err != nil {
		return []domain.AcceptedBidRes{}, err
	}

	return bids, nil
}
