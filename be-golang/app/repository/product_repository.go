package repository

import (
	"be-golang/app/entity"
	"be-golang/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// "encoding/json"
	// "fmt"
	"gorm.io/gorm"
	// "io/ioutil"
	// "net/http"
)

type ExchangeRateResponse struct {
	Bank     string  `json:"bank"`
	Rate     float64 `json:"jual"`
	MataUang string  `json:"matauang"`
}

type ProductRepository interface {
	GetProducts(limit int, offsite int) (map[string]interface{}, error)
	GetAggregatedProducts() ([]map[string]interface{}, error)
}

type ProductRepositoryImpl struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewProductRepository(db *gorm.DB, cfg *config.Config) ProductRepository {
	return &ProductRepositoryImpl{db: db, cfg: cfg}
}

func (r *ProductRepositoryImpl) getExchangeRate() (float64, error) {
	token := r.cfg.EXCHANGE_API_TOKEN
	apiURL := fmt.Sprintf("https://api.kurs.web.id/api/v1?token=%s&bank=bca&matauang=usd", token)

	// Make the HTTP request
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch exchange rate: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var exchangeRateResponse ExchangeRateResponse
	if err := json.Unmarshal(body, &exchangeRateResponse); err != nil {
		return 0, fmt.Errorf("failed to parse exchange rate response: %w", err)
	}

	return exchangeRateResponse.Rate, nil
}

func (r *ProductRepositoryImpl) GetProducts(limit int, offset int) (map[string]interface{}, error) {
	var products []entity.Product
	var total int64

	// ambil kurs dollar
	exchangeRate, err := r.getExchangeRate()
	if err != nil {
		return nil, err
	}

	err = r.db.Offset(offset).Limit(limit).Find(&products).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Model(&entity.Product{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	// buat pagination
	totalPages := (total + int64(limit) - 1) / int64(limit)
	currentPage := offset/limit + 1
	nextPage := currentPage + 1
	if int64(nextPage) > int64(totalPages) {
		nextPage = 0
	}
	previousPage := currentPage - 1
	if previousPage < 1 {
		previousPage = 0
	}

	var result []map[string]interface{}
	for _, product := range products {
		priceUSD := 0.0
		if exchangeRate > 0 {
			priceUSD = product.Price / exchangeRate
		}
		result = append(result, map[string]interface{}{
			"id":         product.ID,
			"product":    product.Product,
			"department": product.Department,
			"priceIDR":   product.Price,
			"priceUSD":   priceUSD,
		})
	}

	pagination := map[string]interface{}{
		"current_page": currentPage,
		"next":         nextPage,
		"per_page":     limit,
		"previous":     previousPage,
		"total_items":  total,
		"total_pages":  totalPages,
	}

	return map[string]interface{}{
		"pagination": pagination,
		"data":       result,
	}, nil
}

func (r *ProductRepositoryImpl) GetAggregatedProducts() ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	// Query untuk aggregat dan sortir
	err := r.db.Model(&entity.Product{}).
		Select("department, product, price").
		Order("price ASC").
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}
