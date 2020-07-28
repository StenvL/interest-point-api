package models

// Point model
type Point struct {
	ID          uint64     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Type        *PointType `json:"type"`
	City        *City      `json:"city"`
	Lon         float32    `json:"lon"`
	Lat         float32    `json:"lat"`
}
