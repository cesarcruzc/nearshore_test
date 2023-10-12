package device

import (
	"fmt"
	"github.com/cesarcruzc/nearshore_test/internal/domain"
	"gorm.io/gorm"
	"log"
	"strings"
)

type (
	Repository interface {
		Create(device *domain.Device) error
		GetAll(filter Filters, offset, limit int64) ([]*domain.Device, error)
		Get(id string) (*domain.Device, error)
		Delete(id string) error
		Update(id string, name *string) error
		Count(filter Filters) (int64, error)
	}

	repository struct {
		log *log.Logger
		db  *gorm.DB
	}
)

func NewRepository(log *log.Logger, db *gorm.DB) Repository {
	return &repository{
		log: log,
		db:  db,
	}
}

func (r *repository) Create(device *domain.Device) error {
	if err := r.db.Create(&device).Error; err != nil {
		r.log.Println("Error creating device: ", err)
		return err
	}

	r.log.Println("Device created with ID: ", device.ID)
	return nil
}

func (r *repository) GetAll(filter Filters, offset, limit int64) ([]*domain.Device, error) {
	var devices []*domain.Device

	tx := r.db.Model(&devices)
	tx = applyFilters(tx, filter)
	tx = tx.Offset(int(offset)).Limit(int(limit))

	if err := tx.Order("created_at desc").Find(&devices).Error; err != nil {
		r.log.Println("Error getting all devices: ", err)
		return nil, err
	}

	return devices, nil
}

func (r *repository) Get(id string) (*domain.Device, error) {
	var device domain.Device

	if err := r.db.Where("id = ?", id).First(&device).Error; err != nil {
		r.log.Println("Error getting device: ", err)
		return nil, err
	}

	r.log.Println("Device found with ID: ", device.ID)
	return &device, nil
}

func (r *repository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&domain.Device{}).Error; err != nil {
		r.log.Println("Error deleting device: ", err)
		return err
	}

	return nil
}

func (r *repository) Update(id string, name *string) error {
	values := make(map[string]interface{})

	if name != nil {
		values["name"] = *name
	}

	if err := r.db.Model(&domain.Device{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {
	if filters.Name != "" {
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		tx = tx.Where("LOWER(name) LIKE ?", filters.Name)
	}

	return tx
}

func (r *repository) Count(filter Filters) (int64, error) {
	var count int64

	tx := r.db.Model(&domain.Device{})
	tx = applyFilters(tx, filter)

	if err := tx.Count(&count).Error; err != nil {
		r.log.Println("Error getting count: ", err)
		return 0, err
	}

	return count, nil
}
