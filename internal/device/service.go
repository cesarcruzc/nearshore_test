package device

import (
	"github.com/cesarcruzc/nearshore_test/internal/domain"
	"log"
)

type (
	Service interface {
		Create(name string) (*domain.Device, error)
		GetAll(filters Filters, offset, limit int64) ([]*domain.Device, error)
		Get(id string) (*domain.Device, error)
		Delete(id string) error
		Update(id string, name *string) error
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

func (s *service) Create(name string) (*domain.Device, error) {
	s.log.Println("Creating device")

	device := &domain.Device{
		Name: name,
	}

	if err := s.repository.Create(device); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return device, nil
}

func (s *service) GetAll(filters Filters, offset, limit int64) ([]*domain.Device, error) {
	s.log.Println("Getting all devices")

	devices, err := s.repository.GetAll(filters, offset, limit)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	return devices, nil
}

func (s *service) Get(id string) (*domain.Device, error) {
	s.log.Println("Getting device")

	device, err := s.repository.Get(id)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	return device, nil
}

func (s *service) Delete(id string) error {
	s.log.Println("Deleting device")

	err := s.repository.Delete(id)
	if err != nil {
		s.log.Println(err)
		return err
	}

	return nil
}

func (s *service) Update(id string, name *string) error {
	s.log.Println("Updating device")
	err := s.repository.Update(id, name)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Count(filter Filters) (int64, error) {
	s.log.Println("Counting devices")

	count, err := s.repository.Count(filter)
	if err != nil {
		s.log.Println(err)
		return 0, err
	}

	return count, nil
}
