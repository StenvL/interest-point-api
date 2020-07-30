package responses

import (
	"github.com/StenvL/interest-points-api/models/domain"
)

//PointDistanceResponse response model
type PointDistanceResponse struct {
	ID          uint64             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Type        *PointTypeResponse `json:"type"`
	City        *CityResponse      `json:"city"`
	Lon         float32            `json:"lon"`
	Lat         float32            `json:"lat"`
	Distance    uint64             `json:"distance"`
	WalkTime    float32            `json:"walkTime"`
}

//NewEmptyPointDistanceResponse creates empty point distance response
func NewEmptyPointDistanceResponse() *PointDistanceResponse {
	return &PointDistanceResponse{
		Type: &PointTypeResponse{},
		City: &CityResponse{},
	}
}

//NewPointDistanceResponse creates point distance response by point domain entity
func NewPointDistanceResponse(p *domain.Point) *PointDistanceResponse {
	const walkSpeed, minutesInHour = 5, 60

	return &PointDistanceResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Type:        NewPointTypeResponse(p.Type),
		City:        NewCityResponse(p.City),
		Lon:         p.Lon,
		Lat:         p.Lat,
		Distance:    p.Distance,
		WalkTime:    float32(p.Distance) / float32(walkSpeed) * minutesInHour,
	}
}
