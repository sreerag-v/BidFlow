package providerRepo

import (
	"fmt"

	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/repository/provider/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"gorm.io/gorm"
)

type ProWorkRepo struct {
	DB *gorm.DB
}

func NewProWorkRepo(db *gorm.DB) interfaces.ProWorkRepo {
	return &ProWorkRepo{
		DB: db,
	}
}

func (p *ProWorkRepo) GetLeadByServiceAndLocation(service, location int) ([]int, error) {
	var id []int64
	if err := p.DB.Raw("SELECT id FROM works WHERE target_profession_id = $1 AND district_id = $2 AND work_status = $3", service, location, "listed").Scan(&id).Error; err != nil {
		return []int{}, err
	}
	var result []int
	for _, v := range id {
		result = append(result, int(v))
	}
	return result, nil
}

func (p *ProWorkRepo) GetDetailsOfAWork(id int, page models.PageNation) (models.MinWorkDetails, error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	var model models.MinWorkDetails

	// Assuming p.DB is an instance of *gorm.DB
	err := p.DB.Table("works").
		Order("id asc").
		Limit(int(limit)).
		Offset(int(offset)).
		Select("works.id, works.street, districts.district as district, states.state as state, professions.profession as profession, users.name AS user, works.work_status").
		Joins("JOIN districts ON districts.id=works.district_id").
		Joins("JOIN states ON states.id=works.state_id").
		Joins("JOIN professions ON professions.id=works.target_profession_id").
		Joins("JOIN users ON users.id=works.user_id").
		Where("works.id = ?", id).
		Scan(&model).Error

	if err != nil {
		return models.MinWorkDetails{}, err
	}

	fmt.Println("model", model)

	return model, nil
}

func (p *ProWorkRepo) GetImagesOfAWork(id int) ([]string, error) {
	var images []string
	if err := p.DB.Raw("SELECT image FROM workspace_images WHERE work_id = $1", id).Scan(&images).Error; err != nil {
		return []string{}, err
	}

	return images, nil
}

func (p *ProWorkRepo) CheckIfAlreadyBidded(work_id, pro_id int) (bool, error) {
	var count int64
	if err := p.DB.Raw(`SELECT COUNT(*) FROM bids WHERE pro_id = $1 AND work_id = $2`, pro_id, work_id).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (p *ProWorkRepo) PlaceBid(model models.PlaceBid, Uid int) error {
	if err := p.DB.Exec(`INSERT INTO bids(work_id,pro_id,estimate,description,user_id) VALUES($1,$2,$3,$4,$5)`, model.WorkID, model.ProID, model.Estimate, model.Description, Uid).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProWorkRepo) ReplaceBidWithNewBid(model models.PlaceBid, Uid int) error {
	if err := p.DB.Exec(`UPDATE bids SET estimate = $1, description = $2 WHERE pro_id = $3 AND work_id = $4 AND user_id = $5`, model.Estimate, model.Description, model.ProID, model.WorkID, Uid).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProWorkRepo) CheckBidExistOrNot(work_id int)(domain.Bid,error){
	var body domain.Bid

	err:=p.DB.Table("bids").Where("work_id = ?",work_id).Scan(&body).Error
	if err!=nil{
		return domain.Bid{},err
	}

	return body,nil
}


func (p *ProWorkRepo) GetAllOtherBidsOnTheLeads(workID int) ([]models.BidDetails, error) {
	var model []models.BidDetails

	if err := p.DB.Table("bids").
		Select("work_id, providers.name AS provider, providers.id AS provider_id, estimate, description").
		Joins("JOIN providers ON bids.pro_id = providers.id").
		Where("bids.work_id = ? AND is_deleted = ?", workID, false).
		Scan(&model).Error; err != nil {
		return []models.BidDetails{}, err
	}

	return model, nil
}

func (p *ProWorkRepo) FindProviderName(pro_id int) (string, error) {
	var name string
	if err := p.DB.Raw(`SELECT name FROM providers WHERE id = $1`, pro_id).Scan(&name).Error; err != nil {
		return "", err
	}

	return name, nil
}

func (p *ProWorkRepo) FindWorkExistOrNot(work_id int) (domain.Work, error) {
	var body domain.Work

	err := p.DB.Table("works").Where("id = ?", work_id).Scan(&body).Error
	if err != nil {
		return domain.Work{}, err
	}

	return body, err
}

func (p *ProWorkRepo) GetAllWorksOfAProvider(pro_id int) ([]int, error) {
	var works []int
	if err := p.DB.Raw(`SELECT id FROM works WHERE pro_id = $1`, pro_id).Scan(&works).Error; err != nil {
		return []int{}, err
	}

	return works, nil
}

func (p *ProWorkRepo) GetCommittedWorksOfAProvider(pro_id int) ([]int, error) {
	var works []int
	if err := p.DB.Raw(`SELECT id FROM works WHERE pro_id = $1 AND work_status = $2`, pro_id, "committed").Scan(&works).Error; err != nil {
		return []int{}, err
	}

	return works, nil
}

func (p *ProWorkRepo) GetCompletedWorksOfAProvider(pro_id int) ([]int, error) {
	var works []int
	if err := p.DB.Raw(`SELECT id FROM works WHERE pro_id = $1 AND work_status = $2`, pro_id, "completed").Scan(&works).Error; err != nil {
		return []int{}, err
	}

	return works, nil
}

func (p *ProWorkRepo) GetAllAcceptedBids(pro_id int, page models.PageNation) ([]domain.AcceptedBidRes, error) {

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	var body []domain.AcceptedBidRes

	err := p.DB.Table("bids").
		Order("id asc").
		Limit(int(limit)).
		Offset(int(offset)).
		Where("pro_id = ? AND accepted_bid = ? AND is_deleted = ?", pro_id, true, true).
		Scan(&body).Error
	if err != nil {
		return []domain.AcceptedBidRes{}, err
	}

	return body, nil
}


func (p *ProWorkRepo) FindUserByUid(Uid int)(domain.User,error){
	var Body domain.User
	err:=p.DB.Table("users").Where("id = ?",Uid).Scan(&Body).Error
	if err!=nil{
		return domain.User{},err
	}
	return Body,nil
}
