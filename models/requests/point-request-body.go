package requests

import "errors"

// PointRequestBody request body for creating and editing point.
type PointRequestBody struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	TypeID      *uint64  `json:"typeId"`
	CityID      *uint64  `json:"cityId"`
	Lon         *float32 `json:"lon"`
	Lat         *float32 `json:"lat"`
}

// IsValid checks if PointRequestBody is valid.
func (p *PointRequestBody) IsValid() error {
	if len(p.Name) == 0 {
		return errors.New("name parameter must be present")
	}

	if p.TypeID == nil {
		return errors.New("typeId parameter must be present")
	}

	if p.CityID == nil {
		return errors.New("cityId parameter must be present")
	}

	if p.Lon == nil || p.Lat == nil {
		return errors.New("lon and lat parameters must be present")
	}

	return nil
}
