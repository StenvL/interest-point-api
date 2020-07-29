package responses

import "github.com/StenvL/interest-points-api/models/domain"

//PointResponse response model
type PointResponse struct {
	ID          uint64             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Type        *PointTypeResponse `json:"type"`
	City        *CityResponse      `json:"city"`
	Lon         float32            `json:"lon"`
	Lat         float32            `json:"lat"`
}

//NewPointResponse creates point response from domain model
func NewPointResponse(point *domain.Point) *PointResponse {
	return &PointResponse{
		ID:          point.ID,
		Name:        point.Name,
		Description: point.Description,
		Type:        NewPointTypeResponse(point.Type),
		City:        NewCityResponse(point.City),
		Lon:         point.Lon,
		Lat:         point.Lat,
	}
}
