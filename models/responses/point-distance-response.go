package responses

//PointDistanceResponse response model
type PointDistanceResponse struct {
	ID          uint64             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Type        *PointTypeResponse `json:"type"`
	City        *CityResponse      `json:"city"`
	Lon         float32            `json:"lon"`
	Lat         float32            `json:"lat"`
	Distance    uint64             `json:"distance"`
	WalkTime    float32            `json:"walkTime"`
}

//NewPointDistanceResponse creates empty point distance response
func NewPointDistanceResponse() *PointDistanceResponse {
	return &PointDistanceResponse{
		Type: &PointTypeResponse{},
		City: &CityResponse{},
	}
}
