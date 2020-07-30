package queries

import "github.com/StenvL/interest-points-api/models/requests"

//PaginationQuery struct for getting paginated data
type PaginationQuery struct {
	Limit  uint64
	Offset uint64
}

//NewPaginationQuery creates PaginationQuery entity from request entity
func NewPaginationQuery(r *requests.PaginationRequest) PaginationQuery {
	return PaginationQuery{
		Limit:  r.Limit,
		Offset: r.Offset,
	}
}
