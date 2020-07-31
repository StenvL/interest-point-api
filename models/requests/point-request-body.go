package requests

// PointRequestBody request body for creating and editing point.
type PointRequestBody struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	TypeID      uint64  `json:"typeId"`
	CityID      uint64  `json:"cityId"`
	Lon         float32 `json:"lon"`
	Lat         float32 `json:"lat"`
}
