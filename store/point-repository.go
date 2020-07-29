package store

import (
	"database/sql"

	"github.com/StenvL/interest-points-api/models/domain"
)

//PointRepository repository for working with points
type PointRepository struct {
	store *Store
}


//GetByID returns the point by its identifier
func (r *PointRepository) GetByID(id uint64) (*domain.Point, error) {
	point := &domain.Point{
		ID:   id,
		Type: &domain.PointType{},
		City: &domain.City{},
	}

	err := r.store.db.QueryRow(
		"SELECT p.name, p.description, p.lon, p.lat, pt.id, pt.name, c.id, c.name FROM point p JOIN point_type pt ON p.type_id = pt.id JOIN city c ON p.city_id = c.id WHERE p.id = ?;",
		id,
	).Scan(&point.Name, &point.Description, &point.Lon, &point.Lat, &point.Type.ID, &point.Type.Name, &point.City.ID, &point.City.Name)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return point, nil
}
