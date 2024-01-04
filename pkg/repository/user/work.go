package userRepo

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/repository/user/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"gorm.io/gorm"
)

type WorkRepo struct {
	DB *gorm.DB
}

func NewWorkRepo(db *gorm.DB) interfaces.WorkRepo {
	return &WorkRepo{
		DB: db,
	}
}

func (work *WorkRepo) ListNewOpening(model domain.ReqWork) error {

	err := work.DB.Exec("INSERT INTO works(street,district_id,state_id,target_profession_id,user_id)VALUES($1,$2,$3,$4,$5)", model.Street, model.DistrictID, model.StateID, model.TargetProfessionID, model.UserID).Error
	if err != nil {
		return err
	}

	return nil
}

func (work *WorkRepo) GetAllWorksOfAUser(id int) ([]int, error) {
	var works []int
	if err := work.DB.Raw(`SELECT id FROM works WHERE user_id = $1`, id).Scan(&works).Error; err != nil {
		return []int{}, err
	}

	return works, nil
}

func (work *WorkRepo) GetDetailsOfAWork(id int) (models.MinWorkDetails, error) {
	var model models.MinWorkDetails
	if err := work.DB.Raw(`SELECT works.id,works.street,districts.district as district,states.state as state,professions.profession as profession,users.name AS user,works.work_status 
	FROM works 
	JOIN districts ON districts.id=works.district_id 
	JOIN states ON states.id=works.state_id 
	JOIN professions ON professions.id=works.target_profession_id 
	JOIN users ON users.id=works.user_id 
	WHERE works.id=$1`, id).Scan(&model).Error; err != nil {
		return models.MinWorkDetails{}, err
	}

	fmt.Println("model", model)

	return model, nil
}
func (work *WorkRepo) GetImagesOfAWork(id int) ([]string, error) {
	var images []string
	if err := work.DB.Raw("SELECT image FROM workspace_images WHERE work_id = $1", id).Scan(&images).Error; err != nil {
		return []string{}, err
	}

	return images, nil
}

func (work *WorkRepo) FindProviderIdFromWork(id int) (int, error) {
	var pro_id int
	if err := work.DB.Raw("SELECT pro_id FROM works WHERE id = $1", id).Scan(&pro_id).Error; err != nil {
		return 0, nil
	}

	return pro_id, nil
}

func (work *WorkRepo) GetAllBids(page models.PageNation, Uid int) ([]models.BidDetails, error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	var model []models.BidDetails

	if err := work.DB.Table("bids").
		Order("bids.id asc").
		Limit(int(limit)).
		Offset(int(offset)).
		Select("work_id, providers.name AS provider, providers.id AS provider_id, estimate, description").
		Joins("JOIN providers ON bids.pro_id = providers.id").
		Where("bids.user_id = ? AND is_deleted = ?", Uid, false).
		Scan(&model).Error; err != nil {
		return []models.BidDetails{}, err
	}

	return model, nil
}

func (work *WorkRepo) GetAllAcceptedBids(page models.PageNation, Uid int) ([]models.BidDetails, error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	var model []models.BidDetails

	if err := work.DB.Table("bids").
		Order("bids.id asc").
		Limit(int(limit)).
		Offset(int(offset)).
		Select("work_id, providers.name AS provider, providers.id AS provider_id, estimate, description").
		Joins("JOIN providers ON bids.pro_id = providers.id").
		Where("bids.user_id = ? AND accepted_bid = ?", Uid, true).
		Scan(&model).Error; err != nil {
		return []models.BidDetails{}, err
	}

	return model, nil
}

func (work *WorkRepo) FindProviderName(pro_id int) (string, error) {
	var name string
	if err := work.DB.Raw("SELECT name FROM providers WHERE id = $1", pro_id).Scan(&name).Error; err != nil {
		return "", err
	}

	return name, nil
}

func (work *WorkRepo) AddImageOfWork(image string, work_id int) error {
	addimage := domain.WorkspaceImages{
		WorkID: work_id,
		Image:  image,
	}

	err := work.DB.Create(&addimage).Error
	if err != nil {
		return err
	}
	return nil
}

func (work *WorkRepo) GetAllCompletedWorksOfAUser(id int) ([]int, error) {
	var works []int
	if err := work.DB.Raw(`SELECT id FROM works WHERE user_id = $1 AND work_status = 'completed' `, id).Scan(&works).Error; err != nil {
		return []int{}, err
	}

	return works, nil
}

func (work *WorkRepo) GetAllOngoingWorksOfAUser(id int) ([]int, error) {
	var works []int
	if err := work.DB.Raw(`SELECT id FROM works WHERE user_id = $1 AND work_status = 'committed' `, id).Scan(&works).Error; err != nil {
		return []int{}, err
	}

	return works, nil
}

func (w *WorkRepo) AssignWorkToProvider(work_id, pro_id int) error {

	err := w.DB.Exec("UPDATE  works SET pro_id = $1,work_status = 'committed' WHERE id = $2", pro_id, work_id).Error
	if err != nil {
		return err
	}

	return nil
}

func (w *WorkRepo) CheckWorkCommitOrNot(id int) (domain.Work, error) {
	var body domain.Work
	if err := w.DB.Table("works").Where("id = ?", id).Scan(&body).Error; err != nil {
		return domain.Work{}, err
	}

	return body, nil
}

func (w *WorkRepo) MakeWorkAsCompleted(id int) error {
	err := w.DB.Exec("UPDATE  works SET work_status = 'completed' WHERE id = $1", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (w *WorkRepo) RateWork(model models.RatingModel, id int) error {

	err := w.DB.Exec("INSERT INTO ratings(rating,feedback,work_id) VALUES ($1,$2,$3)", model.Rating, model.Feedback, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (w *WorkRepo) FindProviderById(id int) (domain.Provider, error) {
	var body domain.Provider

	err := w.DB.Table("providers").Where("id = ?", id).Scan(&body).Error

	if err != nil {
		return domain.Provider{}, err
	}

	return body, nil
}

func (w *WorkRepo) FindBidExistOrNot(pro_id int, bid_id int) (domain.Bid, error) {
	var body domain.Bid

	err := w.DB.Table("bids").Where("id = ? AND pro_id = ?", bid_id, pro_id).Scan(&body).Error
	if err != nil {
		return domain.Bid{}, err
	}

	return body, nil
}

func (w *WorkRepo) AcceptBid(workID int, proID int) error {
	err := w.DB.Table("bids").
		Where("work_id = ? AND pro_id = ?", workID, proID).
		Updates(map[string]interface{}{"accepted_bid": true}).Error

	if err != nil {
		return err
	}

	return nil
}

func (w *WorkRepo) AddAmountInWork(work_id int, Uid int, amount float64) error {
	err := w.DB.Table("works").Where("id = ? AND user_id = ?", work_id, Uid).
		Updates(map[string]interface{}{"bidded_price": amount}).Error

	if err != nil {
		return err
	}

	return nil
}

func (w *WorkRepo) GetAmountFromWork(workID int, userID int) (float64, error) {
	err := w.DB.Table("works").Where("id = ? AND user_id = ?", workID, userID).
		Update("payment_status", false).Error
	if err != nil {
		return 0, err
	}
	var amount float64
	err = w.DB.Table("works").
		Where("id = ? AND user_id = ? AND payment_status = ?", workID, userID, false).
		Select("bidded_price").
		Scan(&amount).Error

	if err != nil {
		return 0, err
	}

	return amount, nil
}

func (w *WorkRepo) UpdateWorkPaymentField(workID int, userID int) error {
	err := w.DB.Table("works").Where("id = ? AND user_id = ?", workID, userID).
		Updates(map[string]interface{}{"payment_status": true}).Error

	if err != nil {
		return err
	}

	return nil
}

func (w *WorkRepo) DeleteBids(workID int, proID int, bidID int) error {
	err := w.DB.Table("bids").
		Where("id = ? AND work_id = ? AND pro_id = ?", bidID, workID, proID).
		Updates(map[string]interface{}{"is_deleted": true}).Error

	if err != nil {
		return err
	}

	return nil
}

func (w *WorkRepo) RazorPaySucess(Uid int, Oid string, Pid string, Sig string, total string) error {
	Rpay := domain.RazorPay{
		UserID:          Uid,
		RazorPaymentId:  Pid,
		Signature:       Sig,
		RazorPayOrderID: Oid,
		AmountPaid:      total,
	}
	err := w.DB.Create(&Rpay).Error

	if err != nil {
		return err
	}

	todyDate := time.Now()
	method := "Razor Pay"
	status := "pending"

	totalprice, err := strconv.Atoi(total)
	if err != nil {
		return errors.New("error in string convertion")
	}

	paymentdata := domain.Payment{
		UserId:        Uid,
		PaymentMethod: method,
		Status:        status,
		Date:          todyDate,
		Totalamount:   totalprice,
	}

	err = w.DB.Create(&paymentdata).Error

	if err != nil {
		return err
	}

	return nil

}
