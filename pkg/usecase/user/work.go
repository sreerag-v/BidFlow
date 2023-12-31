package userUsecase

import (
	"errors"

	"github.com/razorpay/razorpay-go"
	"github.com/sreerag_v/BidFlow/pkg/config"
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

func (work *WorkUsecase) AddImageOfWork(image string, work_id int) error {
	exist, err := work.Repo.GetDetailsOfAWork(work_id)
	if err != nil {
		return err
	}
	if exist.ID == 0 {
		return errors.New("work does not exist")
	}

	return work.Repo.AddImageOfWork(image, work_id)
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

func (work *WorkUsecase) GetAllBids(page models.PageNation, Uid int) ([]models.BidDetails, error) {
	bids, err := work.Repo.GetAllBids(page, Uid)
	if err != nil {
		return []models.BidDetails{}, err
	}

	return bids, nil
}

func (work *WorkUsecase) GetAllAcceptedBids(page models.PageNation,Uid int)([]models.BidDetails,error){
	bids, err := work.Repo.GetAllAcceptedBids(page, Uid)
	if err != nil {
		return []models.BidDetails{}, err
	}

	return bids, nil
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

func (w *WorkUsecase) AssignWorkToProvider(work_id, pro_id int) error {
	commited, err := w.Repo.CheckWorkCommitOrNot(work_id)
	if err != nil {
		return err
	}

	if commited.ID == 0 {
		return errors.New("Work Not Found")
	}

	if commited.WorkStatus == "committed" {
		return errors.New("Work is Already Commited")
	} else if commited.WorkStatus == "completed" {
		return errors.New("Work Already Completed")
	}
	err = w.Repo.AssignWorkToProvider(work_id, pro_id)
	if err != nil {
		return err
	}

	return nil
}

func (w *WorkUsecase) MakeWorkAsCompleted(id int) error {
	committed, err := w.Repo.CheckWorkCommitOrNot(id)
	if err != nil {
		return err
	}

	if committed.ID == 0 {
		return errors.New("Work Not Found")
	}

	if committed.WorkStatus != "committed" {
		return errors.New("Work is not Commited")
	}
	//pass to repository
	err = w.Repo.MakeWorkAsCompleted(id)
	if err != nil {
		return err
	}

	return nil
}

func (w *WorkUsecase) RateWork(model models.RatingModel, work_id int) error {
	exist, err := w.Repo.GetDetailsOfAWork(work_id)
	if err != nil {
		return err
	}

	if exist.ID == 0 {
		return errors.New("Work Not Found")
	}
	//pass to repository
	err = w.Repo.RateWork(model, work_id)
	if err != nil {
		return err
	}

	return nil
}

func (w *WorkUsecase) AcceptBid(work_id int, Pro_id int, bid_id int, Uid int) error {
	exist, err := w.Repo.GetDetailsOfAWork(work_id)
	if err != nil {
		return err
	}

	if exist.ID == 0 {
		return errors.New("Work Not Found")
	}

	get, err := w.Repo.FindBidExistOrNot(Pro_id, bid_id)
	if err != nil {
		return err
	}

	if get.IsDeleted {
		return errors.New("This Bid Already Accepted")
	}

	if get.ID == 0 || get.ProID == 0 {
		return errors.New("bid not found in this id")
	}

	// after that add that amount into the work struct
	err = w.Repo.AddAmountInWork(work_id, Uid, get.Estimate)
	if err != nil {
		return err
	}

	find, err := w.Repo.FindProviderById(Pro_id)
	if err != nil {
		return err
	}

	if find.ID == 0 {
		return errors.New("Provider not found in this id")
	}

	err = w.Repo.AcceptBid(work_id, Pro_id)
	if err != nil {
		return err
	}

	// after accepting bid Delete the bid incoming bids
	err = w.Repo.DeleteBids(work_id, Pro_id, bid_id)
	if err != nil {
		return err
	}

	return nil
}

func (w *WorkUsecase) RazorPaySent(work_id int, Uid int) (interface{}, float64, error) {
	check, err := w.Repo.GetDetailsOfAWork(work_id)
	if err != nil {
		return nil, 0, err
	}

	if check.WorkStatus != "completed" {
		return nil, 0, errors.New("Work is not completed")
	} else if check.WorkStatus == "" {
		return nil, 0, errors.New("Work Not Found")
	}

	amount, err := w.Repo.GetAmountFromWork(work_id, Uid)
	if err != nil {
		return nil, 0, err
	}

	client := razorpay.NewClient(config.GetRazor().Key, config.GetRazor().Secret)
	data := map[string]interface{}{
		"amount":   amount,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)

	if err != nil {
		return nil, 0, err
	}

	// make sure the work table payment was done !!!!
	err = w.Repo.UpdateWorkPaymentField(work_id, Uid)
	if err != nil {
		return nil, 0, err
	}

	value := body["id"]

	return value, amount, nil
}

func (w *WorkUsecase) RazorPaySucess(Uid int, Oid string, Pid string, Sig string, total string) error {
	return w.Repo.RazorPaySucess(Uid, Oid, Pid, Sig, total)
}
