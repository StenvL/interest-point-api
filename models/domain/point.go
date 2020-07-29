package domain

import "github.com/StenvL/interest-points-api/models/requests"

//Point domain model
type Point struct {
	ID          uint64
	Name        string
	Description string
	Type        *PointType
	City        *City
	Lon         float32
	Lat         float32
}

//NewPoint creates point from request model
func NewPoint(pointRequest *requests.PointRequest) *Point {
	return &Point{
		Name:        pointRequest.Name,
		Description: pointRequest.Description,
		Type:        &PointType{ID: pointRequest.TypeID},
		City:        &City{ID: pointRequest.CityID},
		Lon:         pointRequest.Lon,
		Lat:         pointRequest.Lat,
	}
}
