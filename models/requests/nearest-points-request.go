package requests

import (
	"errors"
	"strconv"
)

// NearestPointsRequest is a struct to make request for getting nearest points.
type NearestPointsRequest struct {
	*PaginationRequest
	Radius uint64
	Lon    float64
	Lat    float64
}

// NewNearestPointsRequest creates NearestPointsRequest struct by string params.
func NewNearestPointsRequest(lonParam string, latParam string, radiusParam string, limit string, offset string) (*NearestPointsRequest, error) {
	if len(lonParam) == 0 || len(latParam) == 0 {
		return nil, errors.New("Coordinates must be present")
	}

	lon, err := strconv.ParseFloat(lonParam, 32)
	if err != nil {
		return nil, errors.New("Lon param is incorrect: " + err.Error())
	}

	lat, err := strconv.ParseFloat(latParam, 32)
	if err != nil {
		return nil, errors.New("Lat param is incorrect: " + err.Error())
	}

	if lon < -90 || lon > 90 || lat < -90 && lat > 90 {
		return nil, errors.New("Coordinate values are out of range")
	}

	var radius uint64

	if len(radiusParam) == 0 {
		radius = 1
	} else {
		radius, err = strconv.ParseUint(radiusParam, 10, 32)
		if err != nil {
			return nil, errors.New("Radius param is incorrect: " + err.Error())
		}
	}

	var p *PaginationRequest

	p, err = NewPaginationRequest(limit, offset, 5, 0)
	if err != nil {
		return nil, errors.New("Failed to create pagination request")
	}

	return &NearestPointsRequest{
		Radius:            radius,
		Lon:               lon,
		Lat:               lat,
		PaginationRequest: p,
	}, nil
}
