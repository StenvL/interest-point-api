package domain

//Point domain model
type Point struct {
	ID          uint64
	Name        string
	Description string
	Type        *PointType
	City        *City
	Lon         float32
	Lat         float32
}
