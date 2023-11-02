package repository

import (
	"be-62test/dto"
	"be-62test/entity"
	"fmt"

	"gorm.io/gorm"
)

type BusinessRepository interface {
	GetBusiness(page, perPage int) ([]entity.Business, int64,error)
	PostBusiness(business dto.Business) (entity.Business, error)
	UpdateBusiness(business entity.Business) (entity.Business, error)
	DeleteBusiness(businessId int) error
	FindById(businessId int) (entity.Business, error)
	SearchBusiness(name, location, categories string,page, perPage int) ([]entity.Business, int64, error)
}

type businessRepository struct {
	db *gorm.DB
}

func NewBusinessRepository(db *gorm.DB) *businessRepository {
	return &businessRepository{db: db}
}

func (r *businessRepository) GetBusiness(page, perPage int) ([]entity.Business, int64,error) {
	var business []entity.Business
	var total int64
	
	if err := r.db.Find(&business).Error; err != nil {
		return business, 0, err
	}

	query := r.db.Model(&entity.Business{})

	offset := (page - 1) * perPage

	err := query.Offset(offset).Limit(perPage).Find(&business).Error
	if err != nil {
		return business, 0, err
	}

	r.db.Model(&entity.Business{}).Count(&total)

	return business, total, nil
}

func (r *businessRepository) PostBusiness(business dto.Business) (entity.Business, error) {
	newBusiness := entity.Business{
		Name:       business.Name,
		Location:   business.Location,
		Latitude:   business.Latitude,
		Longitude:  business.Longitude,
		Categories: business.Categories,
		ImageLink: business.ImageLink,
	}

	if err := r.db.Create(&newBusiness).Error; err != nil {
		return newBusiness, err
	}

	return newBusiness, nil
}

func (r *businessRepository) UpdateBusiness(updatedBusiness entity.Business) (entity.Business, error) {
	if err := r.db.Save(&updatedBusiness).Error; err != nil {
        return entity.Business{}, err
    }
    return updatedBusiness, nil
}

func (r *businessRepository) DeleteBusiness(businessId int) error {
	if err := r.db.Where("business_id = ?", businessId).Delete(&entity.Business{}).Error; err != nil {
		return fmt.Errorf("business_id not found")
	}

	return nil
}

func (r *businessRepository) FindById(businessId int) (entity.Business, error) {
	var business entity.Business
	if err := r.db.Where("business_id = ?", businessId).First(&business).Error; err != nil {
		return business, fmt.Errorf("business_id not found")
	}
	
	return business, nil
}

func (r *businessRepository) SearchBusiness(name, location, categories string, page, perPage int) ([]entity.Business, int64, error) {
	var business []entity.Business
	var total int64
	
	query := r.db.Model(&entity.Business{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if location != "" {
		query = query.Where("location LIKE ?", "%"+location+"%")
	}

	if categories != "" {
		query = query.Where("categories LIKE ?", "%"+categories+"%")
	}

	offset := (page - 1) * perPage

	err := query.Offset(offset).Limit(perPage).Find(&business).Error
	if err != nil {
		return business, 0, err
	}

	r.db.Model(&entity.Business{}).Count(&total)

	return business, total, nil
}