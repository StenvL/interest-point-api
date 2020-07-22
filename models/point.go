package models

// Point model
type Point struct {
	BaseModel
	Name        string
	Description string
	Type        *PointType
	City        *City
	Lon         float32
	Lat         float32
}
