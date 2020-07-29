package services

import (
	responses "github.com/StenvL/interest-points-api/models/responses"
	"github.com/StenvL/interest-points-api/store"
)

//PointService service that provides methods to work with point entity
type PointService struct {
	store *store.Store
}

//NewPointService creates an instance of PointService
func NewPointService(store *store.Store) *PointService {
	return &PointService{
		store: store,
	}
}

//GetByID returns point by its identifier
func (p *PointService) GetByID(id uint64) (*responses.PointResponse, error) {
	point, err := p.store.Point().GetByID(id)

	if err != nil {
		return nil, err
	}

	pointResponse := responses.NewPointResponse(point)

	return pointResponse, nil
}
