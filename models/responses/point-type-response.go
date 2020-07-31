package responses

import "github.com/StenvL/interest-points-api/models/domain"

// PointTypeResponse response model.
type PointTypeResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// NewPointTypeResponse creates PointTypeResponse struct from domain model.
func NewPointTypeResponse(pointType *domain.PointType) *PointTypeResponse {
	return &PointTypeResponse{
		ID:   pointType.ID,
		Name: pointType.Name,
	}
}
