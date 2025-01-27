package service

import (
	"be-golang/app/repository"
	"be-golang/config"
)

type ProductService interface {
	GetProducts(limit int, offsite int) (map[string]interface{}, error)
	GetAggregatedProducts() ([]map[string]interface{}, error)
}

type ProductServiceImpl struct {
	repo repository.ProductRepository
	cfg  *config.Config
}

func NewProductService(repo repository.ProductRepository, cfg *config.Config) ProductService {
	return &ProductServiceImpl{repo: repo, cfg: cfg}
}

func (s *ProductServiceImpl) GetProducts(limit int, offsite int) (map[string]interface{}, error) {
	return s.repo.GetProducts(limit, offsite)
}
func (s *ProductServiceImpl) GetAggregatedProducts() ([]map[string]interface{}, error) {
	return s.repo.GetAggregatedProducts()
}
