package providerRepo

import (
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"github.com/sreerag_v/BidFlow/pkg/repository/provider/interfaces"
	"github.com/sreerag_v/BidFlow/pkg/utils/models"
	"gorm.io/gorm"
)

type ProviderRepo struct {
	DB *gorm.DB
}

func NewProviderRepo(db *gorm.DB) interfaces.ProviderRepo {
	return &ProviderRepo{
		DB: db,
	}
}

func (pro *ProviderRepo) CheckPhoneNumber(num string) (bool, error) {
	var count int64
	if err := pro.DB.Raw("SELECT COUNT(*) FROM providers WHERE phone = ?", num).Scan(&count).Error; err != nil {
		return true, err
	}

	return count > 0, nil
}

func (pro *ProviderRepo) Register(model models.ProviderRegister) (int, error) {
	err := pro.DB.Exec("INSERT INTO providers(name,email,password,phone)VALUES($1,$2,$3,$4)", model.Name, model.Email, model.Password, model.Phone).Error
	if err != nil {
		return 0, err
	}

	var id int64
	err = pro.DB.Raw("Select id from providers where phone = $1", model.Phone).Scan(&id).Error
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (pro *ProviderRepo) UploadDoc(id int, Fname string) error {
	if err := pro.DB.Exec("INSERT INTO id_proofs(pro_id,id_proof) VALUES($1,$2)", id, Fname).Error; err != nil {
		return err
	}

	return nil
}

func (pro *ProviderRepo) CheckProExistOrNot(name string) (bool, error) {
	var count int64
	if err := pro.DB.Table("providers").Where("name = ?", name).Count(&count).Error; err != nil {
		return true, err
	}

	// If count is greater than 0, it means a record with the given name exists
	return count > 0, nil
}

func (pro *ProviderRepo) GetProDetails(name string) (domain.Provider, error) {
	var model domain.Provider
	if err := pro.DB.Table("providers").Where("name = ?", name).Scan(&model).Error; err != nil {
		return domain.Provider{}, err
	}

	return model, nil
}
