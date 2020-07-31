package requests

import (
	"errors"
	"strconv"
)

// PaginationRequest is a struct to make request for getting paginated data.
type PaginationRequest struct {
	Limit  uint64
	Offset uint64
}

// NewPaginationRequest creates PaginationRequest struct by string params.
func NewPaginationRequest(limitParam string, offsetParam string, defaultLimit uint64, defaultOffset uint64) (*PaginationRequest, error) {
	var limit, offset uint64
	var err error

	if len(limitParam) == 0 {
		limit = defaultLimit
	} else {
		limit, err = strconv.ParseUint(limitParam, 10, 32)
		if err != nil {
			return nil, errors.New("Limit param is incorrect")
		}
	}

	if len(offsetParam) == 0 {
		offset = defaultOffset
	} else {
		offset, err = strconv.ParseUint(offsetParam, 10, 32)
		if err != nil {
			return nil, errors.New("Offset param is incorrect")
		}
	}

	return &PaginationRequest{
		Limit:  limit,
		Offset: offset,
	}, nil
}
