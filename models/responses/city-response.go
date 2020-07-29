package responses

import "github.com/StenvL/interest-points-api/models/domain"

// CityResponse response model
type CityResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

//NewCityResponse creates point type response from domain model
func NewCityResponse(city *domain.City) *CityResponse {
	return &CityResponse{
		ID:   city.ID,
		Name: city.Name,
	}
}
