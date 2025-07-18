package repository

import (
	"github.com/Mukam21/REST-API_online-marketplace/pkg/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetFiltered(minPrice, maxPrice, page, pageSize int, sortBy, order string) ([]models.Order, error)
}

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{DB: db}
}

func (r *orderRepository) Create(order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *orderRepository) GetFiltered(minPrice, maxPrice, page, pageSize int, sortBy, order string) ([]models.Order, error) {
	var orders []models.Order
	offset := (page - 1) * pageSize

	query := r.DB.Preload("User").Where("price >= ?", minPrice)
	if maxPrice > 0 {
		query = query.Where("price <= ?", maxPrice)
	}

	if sortBy != "" {
		sortStr := sortBy
		if order == "desc" {
			sortStr += " desc"
		} else {
			sortStr += " asc"
		}
		query = query.Order(sortStr)
	} else {
		query = query.Order("created_at desc")
	}

	err := query.Offset(offset).Limit(pageSize).Find(&orders).Error
	return orders, err
}
