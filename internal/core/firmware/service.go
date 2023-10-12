package firmware

import (
	"github.com/cesarcruzc/nearshore_test/internal/core/domain"
	"log"
	"time"
)

type (
	Service interface {
		Create(name, deviceId, version, releaseNotes, releaseDate, url string) (*domain.Firmware, error)
		GetAll(filters Filters, offset, limit int64) ([]*domain.Firmware, error)
		Get(id string) (*domain.Firmware, error)
		Delete(id string) error
		Update(id string, name, deviceId, version, releaseNotes, releaseDate, url *string) error
		Count(filter Filters) (int64, error)
	}

	service struct {
		log        *log.Logger
		repository Repository
	}

	Filters struct {
		Name string
	}
)

func NewService(log *log.Logger, repository Repository) Service {
	return &service{
		log:        log,
		repository: repository,
	}
}

func (s *service) Create(name, deviceId, version, releaseNotes, releaseDate, url string) (*domain.Firmware, error) {
	s.log.Println("Creating firmware")

	releaseDateParsed, err := time.Parse("2006-01-02", releaseDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	firmware := &domain.Firmware{
		Name:         name,
		DeviceID:     deviceId,
		Version:      version,
		ReleaseNotes: releaseNotes,
		ReleaseDate:  releaseDateParsed,
		Url:          url,
	}

	if err := s.repository.Create(firmware); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return firmware, nil
}

func (s *service) GetAll(filters Filters, offset, limit int64) ([]*domain.Firmware, error) {
	s.log.Println("Getting all firmwares")

	firmwares, err := s.repository.GetAll(filters, offset, limit)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	return firmwares, nil
}

func (s *service) Get(id string) (*domain.Firmware, error) {
	s.log.Println("Getting firmware")

	firmware, err := s.repository.Get(id)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	return firmware, nil
}

func (s *service) Delete(id string) error {
	s.log.Println("Deleting firmware")

	err := s.repository.Delete(id)
	if err != nil {
		s.log.Println(err)
		return err
	}

	return nil
}

func (s *service) Update(id string, name, deviceId, version, releaseNotes, releaseDate, url *string) error {
	s.log.Println("Updating firmware")
	err := s.repository.Update(id, name, deviceId, version, releaseNotes, releaseDate, url)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Count(filter Filters) (int64, error) {
	s.log.Println("Counting firmwares")

	count, err := s.repository.Count(filter)
	if err != nil {
		s.log.Println(err)
		return 0, err
	}

	return count, nil
}
