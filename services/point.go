package services

import (
	"github.com/StenvL/interest-points-api/models/domain"
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

//GetAll returns all points
func (p *PointService) GetAll(cityID uint64) ([]*domain.Point, error) {
	points, err := p.store.Point().GetAll(cityID)

	if err != nil {
		return nil, err
	}

	return points, nil
}

//GetByID returns point by its identifier
func (p *PointService) GetByID(id uint64) (*domain.Point, error) {
	point, err := p.store.Point().GetByID(id)

	if err != nil {
		return nil, err
	}

	if point == nil {
		return nil, nil
	}

	return point, nil
}

//Create new point
func (p *PointService) Create(point *domain.Point) error {
	err := p.store.Point().Create(point)

	if err != nil {
		return err
	}

	return nil
}
