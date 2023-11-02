package service

import (
	"be-62test/dto"
	"be-62test/entity"
	"be-62test/repository"
	"fmt"
)

type BusinessService interface {
	GetBusiness(page, perPage int) ([]entity.Business, int64, error)
	PostBusiness(business dto.Business) (entity.Business, error)
	UpdateBusiness(business entity.Business) (entity.Business, error)
	DeleteBusiness(businessId int) error
	FindById(businessId int) (entity.Business, error)
	SearchBusiness(name, location, categories string, page, perPage int) ([]entity.Business, int64, error)
}

type businessService struct {
	repository repository.BusinessRepository
}

func NewBusinessService(repository repository.BusinessRepository) *businessService {
	return &businessService{repository: repository}
}

func (s *businessService) GetBusiness(page, perPage int) ([]entity.Business, int64, error) {
	business, total, err := s.repository.GetBusiness(page, perPage)
	if err != nil {
		return nil, 0, err
	}

	return business, total, nil
}

func (s *businessService) PostBusiness(business dto.Business) (entity.Business, error) {
	return s.repository.PostBusiness(business)
}

func (s *businessService) UpdateBusiness(business entity.Business) (entity.Business, error) {
    existingBusiness, err := s.repository.FindById(business.BusinessId)
    if err != nil {
        return entity.Business{}, fmt.Errorf("business not found")
    }

    existingBusiness.Name = business.Name
    existingBusiness.Location = business.Location
    existingBusiness.Latitude = business.Latitude
    existingBusiness.Longitude = business.Longitude
    existingBusiness.Categories = business.Categories
    existingBusiness.ImageLink = business.ImageLink

    updatedBusiness, err := s.repository.UpdateBusiness(existingBusiness)
    if err != nil {
        return entity.Business{}, fmt.Errorf("failed to update business")
    }

    return updatedBusiness, nil
}


func (s *businessService) DeleteBusiness(businessId int) error {
	return s.repository.DeleteBusiness(businessId)
}

func (s *businessService) FindById(businessId int) (entity.Business, error) {
	return s.repository.FindById(businessId)
}

func (s *businessService) SearchBusiness(name, location, categories string, page, perPage int) ([]entity.Business, int64, error) {
	business, total, err := s.repository.SearchBusiness(name, location, categories, page, perPage)

	if err != nil {
		return nil, 0, err
	}

	return business, total, nil
}