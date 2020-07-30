package services

import (
	"github.com/StenvL/interest-points-api/models/domain"
	"github.com/StenvL/interest-points-api/models/domain/queries"
	"github.com/StenvL/interest-points-api/models/requests"
	"github.com/StenvL/interest-points-api/models/responses"
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

//GetAllByCity returns all points by city
func (p *PointService) GetAllByCity(r *requests.PointsRequest) ([]*responses.PointResponse, error) {
	points, err := p.store.Point().GetAllByCity(queries.NewPointsQuery(r))

	if err != nil {
		return nil, err
	}

	pointsResponse := make([]*responses.PointResponse, 0)
	for _, point := range points {
		pointsResponse = append(pointsResponse, responses.NewPointResponse(point))
	}

	return pointsResponse, nil
}

//GetNearest returns nearest points by radius
func (p *PointService) GetNearest(r *requests.NearestPointsRequest) ([]*responses.PointDistanceResponse, error) {
	points, err := p.store.Point().GetNearest(queries.NewNearestPointsQuery(r))

	if err != nil {
		return nil, err
	}

	pointsResponse := make([]*responses.PointDistanceResponse, 0)
	for _, point := range points {
		pointsResponse = append(pointsResponse, responses.NewPointDistanceResponse(point))
	}

	return pointsResponse, nil
}

//GetByID returns point by its identifier
func (p *PointService) GetByID(id uint64) (*responses.PointResponse, error) {
	point, err := p.store.Point().GetByID(id)

	if err != nil {
		return nil, err
	}

	if point == nil {
		return nil, nil
	}

	return responses.NewPointResponse(point), nil
}

//Create new point
func (p *PointService) Create(r *requests.PointRequestBody) (uint64, error) {
	pointDomain := domain.NewPoint(r)

	err := p.store.Point().Create(pointDomain)

	if err != nil {
		return 0, err
	}

	return pointDomain.ID, nil
}

//Update existent point
func (p *PointService) Update(id uint64, r *requests.PointRequestBody) (*responses.PointResponse, error) {
	pointDomain := domain.NewPoint(r)
	pointDomain.ID = id

	pointDomain, err := p.store.Point().Update(pointDomain)

	if err != nil {
		return nil, err
	}

	return responses.NewPointResponse(pointDomain), nil
}
