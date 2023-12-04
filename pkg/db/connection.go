package db

import (
	"fmt"

	"github.com/sreerag_v/BidFlow/pkg/config"
	"github.com/sreerag_v/BidFlow/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})

	db.AutoMigrate(domain.Admin{})
	db.AutoMigrate(domain.Provider{})
	db.AutoMigrate(domain.IdProof{})
	db.AutoMigrate(domain.User{})
	db.AutoMigrate(domain.Category{})
	db.AutoMigrate(domain.State{})
	db.AutoMigrate(domain.District{})
	db.AutoMigrate(domain.Profession{})
	db.AutoMigrate(domain.Probook{})
	db.AutoMigrate(domain.PreferredLocation{})
	db.AutoMigrate(domain.Rating{})
	db.AutoMigrate(domain.Work{})
	return db, dbErr
}
