package requests

import (
	"errors"
	"strconv"
)

// PointsRequest is a struct to make request for getting points index.
type PointsRequest struct {
	*PaginationRequest
	City uint64
}

// NewPointsRequest creates PointsRequest struct by string params.
func NewPointsRequest(city string, limit string, offset string) (*PointsRequest, error) {
	if len(city) == 0 {
		return nil, errors.New("City param must be present")
	}

	cityID, err := strconv.ParseUint(city, 10, 32)
	if err != nil {
		return nil, errors.New("City param is incorrect")
	}

	var p *PaginationRequest

	p, err = NewPaginationRequest(limit, offset, 50, 0)
	if err != nil {
		return nil, errors.New("Failed to create pagination request")
	}

	return &PointsRequest{
		City:              cityID,
		PaginationRequest: p,
	}, nil
}
