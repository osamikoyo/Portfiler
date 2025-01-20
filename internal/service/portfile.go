package service

import (
	"github.com/osamikoyo/portfiler/internal/data"
	"github.com/osamikoyo/portfiler/internal/data/models"
)

type PortfileService struct{
	Data *data.PortfolioStorage
}

func New() (PortfileService, error){
	data, err := data.New()
	return PortfileService{Data: data}, err
}

func (service *PortfileService) Add(portfile models.Portfolio) error {
	return service.Data.Add(portfile)
}

func (service *PortfileService) Get(title string) ([]models.Portfolio, error) {
	return service.Data.Get(title)
}

func (service *PortfileService) AddReview(title string, review models.Review) error {
	return service.Data.AddReview(title, review)
}