package responses

// PointResponse model
type PointDetailedResponse struct {
	ID          uint64             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Type        *PointTypeResponse `json:"type"`
	City        *CityResponse      `json:"city"`
	Lon         float32            `json:"lon"`
	Lat         float32            `json:"lat"`
	WalkTime    time               `json:"walkTime"`
}
