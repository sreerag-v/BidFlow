package providerRepo

import (
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/repository/provider/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"gorm.io/gorm"
)

type ProfileRepo struct{
	DB *gorm.DB
}

func NewProfileRepo(db *gorm.DB)interfaces.ProfileRepo{
	return &ProfileRepo{
		DB: db,
	}
}

func (pr *ProfileRepo) CheckIfServiceIsAlreadyRegistered(user_id, service_id int) (bool, error) {
	var count int64
	if err := pr.DB.Raw("SELECT COUNT(*) FROM probooks WHERE pro_id = $1 AND profession_id = $2", user_id, service_id).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (pr *ProfileRepo) AddService(user_id, service_id int) error {

	if err := pr.DB.Exec("INSERT INTO probooks(pro_id,profession_id)VALUES($1,$2)", user_id, service_id).Error; err != nil {
		return err
	}

	return nil
}

func (pr *ProfileRepo) DeleteService(user_id, service_id int) error {

	if err := pr.DB.Exec("DELETE FROM probooks WHERE pro_id = $1 AND profession_id = $2", user_id, service_id).Error; err != nil {
		return err
	}

	return nil
}
func (pr *ProfileRepo) CheckIfDistrictIsAlreadyAdded(id, district int) (bool, error) {
	var count int64
	if err := pr.DB.Raw("SELECT COUNT(*) FROM preferred_locations WHERE pro_id = $1 AND district_id = $2", id, district).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (pr *ProfileRepo) AddLocation(id, district int) error {

	if err := pr.DB.Exec("INSERT INTO preferred_locations(pro_id,district_id)VALUES($1,$2)", id, district).Error; err != nil {
		return err
	}

	return nil
}
func (pr *ProfileRepo) RemovePreferredLocation(id, district int) error {

	if err := pr.DB.Exec("DELETE FROM preferred_locations WHERE pro_id = $1 AND district_id = $2", id, district).Error; err != nil {
		return err
	}

	return nil
}

func (pr *ProfileRepo) GetAllServiceIdsOfAProvider(id int) ([]int, error) {

	var services []int

	if err := pr.DB.Raw("SELECT profession_id FROM probooks WHERE pro_id = $1", id).Scan(&services).Error; err != nil {
		return []int{}, err
	}

	return services, nil
}


func (pr *ProfileRepo) FindServiceDetailsFromID(id int) (domain.Profession, error) {

	var service domain.Profession

	if err := pr.DB.Raw("SELECT * FROM professions WHERE id = $1", id).Scan(&service).Error; err != nil {
		return domain.Profession{}, err
	}

	return service, nil
}

func (pr *ProfileRepo) GetAllPreferredLocations(id int) ([]int, error) {

	var locations []int

	if err := pr.DB.Raw("SELECT district_id FROM preferred_locations WHERE pro_id = $1", id).Scan(&locations).Error; err != nil {
		return []int{}, err
	}

	return locations, nil
}

func (pr *ProfileRepo) GetLocationDetails(id int) (models.GetLocations, error) {

	var service models.GetLocations

	if err := pr.DB.Raw("select district,states.state AS state from districts join states on states.id=districts.state_id where districts.id = $1", id).Scan(&service).Error; err != nil {
		return models.GetLocations{}, err
	}

	return service, nil
}

func (pr *ProfileRepo) FindProviderDetails(id int) (domain.Provider, error) {

	var pro domain.Provider
	if err := pr.DB.Raw("SELECT * FROM providers WHERE id = $1", id).Scan(&pro).Error; err != nil {
		return domain.Provider{}, err
	}

	return pro, nil
}

func (pr *ProfileRepo) GetRatingsOfAllRecordsOfAProvider(id int) ([]int, error) {

	var ratings []int

	if err := pr.DB.Raw(`SELECT ratings.rating FROM works JOIN ratings ON works.id = ratings.work_id JOIN providers ON providers.id = works.pro_id WHERE providers.id = $1`, id).Scan(&ratings).Error; err != nil {
		return []int{}, err
	}

	return ratings, nil
}