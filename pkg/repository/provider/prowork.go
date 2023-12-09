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

func (p *ProWorkRepo) PlaceBid(model models.PlaceBid) error {
	if err := p.DB.Exec(`INSERT INTO bids(work_id,pro_id,estimate,description) VALUES($1,$2,$3,$4)`, model.WorkID, model.ProID, model.Estimate, model.Description).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProWorkRepo) ReplaceBidWithNewBid(model models.PlaceBid) error {
	if err := p.DB.Exec(`UPDATE bids SET estimate = $1, description = $2 WHERE pro_id = $3 AND work_id = $4`, model.Estimate, model.Description, model.ProID, model.WorkID).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProWorkRepo) GetAllOtherBidsOnTheLeads(work_id int) ([]models.BidDetails, error) {
	var model []models.BidDetails
	if err := p.DB.Raw(`SELECT bids.id,providers.name AS provider,estimate,description FROM bids JOIN providers ON bids.pro_id = providers.id WHERE bids.work_id = $1`, work_id).Scan(&model).Error; err != nil {
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

func (p *ProWorkRepo) GetAllAcceptedBids(pro_id int, page models.PageNation) ([]domain.Bid, error) {

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	var body []domain.Bid

	err := p.DB.Table("bids").
		Order("id asc").
		Limit(int(limit)).
		Offset(int(offset)).
		Where("pro_id = ? AND accepted_bid = ?", pro_id, true).
		Scan(&body).Error
	if err != nil {
		return []domain.Bid{}, err
	}

	return body, nil
}
