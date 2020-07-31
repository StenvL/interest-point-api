package queries

import (
	"github.com/StenvL/interest-points-api/models/requests"
)

// PointsQuery is a struct for getting points by query.
type PointsQuery struct {
	City uint64
	PaginationQuery
}

// NewPointsQuery creates NewPointsQuery entity from request entity.
func NewPointsQuery(r *requests.PointsRequest) PointsQuery {
	return PointsQuery{
		City:            r.City,
		PaginationQuery: NewPaginationQuery(r.PaginationRequest),
	}
}
