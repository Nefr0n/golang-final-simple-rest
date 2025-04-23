package repository

import (
	"github.com/Skapar/simple-rest/internal/models/entities"
	"gorm.io/gorm"
)

type ProfileRepositoryImpl struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &ProfileRepositoryImpl{db: db}
}

func (r *ProfileRepositoryImpl) GetProfileByID(id int64) (*entities.Profile, error) {
	var profile entities.Profile
	if err := r.db.First(&profile, id).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *ProfileRepositoryImpl) CreateProfile(profile *entities.Profile) (*entities.Profile, error) {
	if err := r.db.Create(profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

func (r *ProfileRepositoryImpl) UpdateProfile(profile *entities.Profile) (*entities.Profile, error) {
	if err := r.db.Save(profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

func (r *ProfileRepositoryImpl) DeleteProfile(id int64) error {
	if err := r.db.Delete(&entities.Profile{}, id).Error; err != nil {
		return err
	}
	return nil
}
