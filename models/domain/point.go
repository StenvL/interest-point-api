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
	Distance    uint64
}

//NewEmptyPoint creates empty point
func NewEmptyPoint() *Point {
	return &Point{
		Type: &PointType{},
		City: &City{},
	}
}

//NewPoint creates point from request model
func NewPoint(pointRequestBody *requests.PointRequestBody) *Point {
	return &Point{
		Name:        pointRequestBody.Name,
		Description: pointRequestBody.Description,
		Type:        &PointType{ID: pointRequestBody.TypeID},
		City:        &City{ID: pointRequestBody.CityID},
		Lon:         pointRequestBody.Lon,
		Lat:         pointRequestBody.Lat,
	}
}
