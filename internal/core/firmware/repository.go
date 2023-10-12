package firmware

import (
	"fmt"
	"github.com/cesarcruzc/nearshore_test/internal/core/domain"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

type (
	Repository interface {
		Create(firmware *domain.Firmware) error
		GetAll(filter Filters, offset, limit int64) ([]*domain.Firmware, error)
		Get(id string) (*domain.Firmware, error)
		Delete(id string) error
		Update(id string, name, deviceId, version, releaseNotes, releaseDate, url *string) error
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

func (r *repository) Create(firmware *domain.Firmware) error {
	if err := r.db.Create(&firmware).Error; err != nil {
		r.log.Println("Error creating firmware: ", err)
		return err
	}

	r.log.Println("Firmware created with ID: ", firmware.ID)
	return nil
}

func (r *repository) GetAll(filter Filters, offset, limit int64) ([]*domain.Firmware, error) {
	var firmwares []*domain.Firmware

	tx := r.db.Model(&firmwares)
	tx = applyFilters(tx, filter)
	tx = tx.Offset(int(offset)).Limit(int(limit))

	if err := tx.Order("created_at desc").Find(&firmwares).Error; err != nil {
		r.log.Println("Error getting all firmwares: ", err)
		return nil, err
	}

	return firmwares, nil
}

func (r *repository) Get(id string) (*domain.Firmware, error) {
	var firmware domain.Firmware

	if err := r.db.Where("id = ?", id).First(&firmware).Error; err != nil {
		r.log.Println("Error getting firmware: ", err)
		return nil, err
	}

	r.log.Println("Firmware found with ID: ", firmware.ID)
	return &firmware, nil
}

func (r *repository) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&domain.Firmware{}).Error; err != nil {
		r.log.Println("Error deleting firmware: ", err)
		return err
	}

	return nil
}

func (r *repository) Update(id string, name, deviceId, version, releaseNotes, releaseDate, url *string) error {
	values := make(map[string]interface{})

	if name != nil {
		values["name"] = *name
	}

	if deviceId != nil {
		values["device_id"] = *deviceId
	}

	if version != nil {
		values["version"] = *version
	}

	if releaseNotes != nil {
		values["release_notes"] = *releaseNotes
	}

	if releaseDate != nil {
		releaseDateParsed, err := time.Parse("2006-01-02", *releaseDate)
		if err != nil {
			r.log.Println(err)
			return err
		}
		values["release_date"] = releaseDateParsed
	}

	if url != nil {
		values["url"] = *url
	}

	if err := r.db.Model(&domain.Firmware{}).Where("id = ?", id).Updates(values).Error; err != nil {
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

	tx := r.db.Model(&domain.Firmware{})
	tx = applyFilters(tx, filter)

	if err := tx.Count(&count).Error; err != nil {
		r.log.Println("Error getting count: ", err)
		return 0, err
	}

	return count, nil
}
