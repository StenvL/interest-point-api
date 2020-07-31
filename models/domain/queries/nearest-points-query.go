package queries

import (
	"github.com/StenvL/interest-points-api/models/requests"
)

// NearestPointsQuery struct for getting nearest points.
type NearestPointsQuery struct {
	Radius uint64
	Lon    float64
	Lat    float64
	PaginationQuery
}

// NewNearestPointsQuery creates NearestPointsQuery entity from request entity.
func NewNearestPointsQuery(r *requests.NearestPointsRequest) NearestPointsQuery {
	return NearestPointsQuery{
		Radius:          r.Radius,
		Lon:             r.Lon,
		Lat:             r.Lat,
		PaginationQuery: NewPaginationQuery(r.PaginationRequest),
	}
}
