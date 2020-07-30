package queries

import (
	"github.com/StenvL/interest-points-api/models/requests"
)

//NearestPointsQuery struct to for getting nearest points
type NearestPointsQuery struct {
	Radius uint64
	Lon    float64
	Lat    float64
}

//NewNearestPointsQuery creates NearestPointsQuery entity from request entity
func NewNearestPointsQuery(r *requests.NearestPointsRequest) NearestPointsQuery {
	return NearestPointsQuery{
		Radius: r.Radius,
		Lon:    r.Lon,
		Lat:    r.Lat,
	}
}
