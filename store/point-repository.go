package store

import (
	"database/sql"
	"fmt"

	"github.com/StenvL/interest-points-api/models/domain"
	"github.com/StenvL/interest-points-api/models/domain/queries"
)

// PointRepository repository for working with points.
type PointRepository struct {
	store *Store
}

// Create creates new point.
func (r *PointRepository) Create(point *domain.Point) error {
	res, err := r.store.db.Exec(
		"INSERT INTO point (name, description, type_id, city_id, lon, lat) VALUES (?, ?, ?, ?, ?, ?);",
		point.Name,
		point.Description,
		point.Type.ID,
		point.City.ID,
		point.Lon,
		point.Lat,
	)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return err
	}

	point.ID = uint64(id)

	return nil
}

// Update updates existent point.
func (r *PointRepository) Update(point *domain.Point) (*domain.Point, error) {
	_, err := r.store.db.Exec(
		"UPDATE point SET name = ?, description = ?, type_id = ?, city_id = ?, lon = ?, lat = ? WHERE ID = ?",
		point.Name,
		point.Description,
		point.Type.ID,
		point.City.ID,
		point.Lon,
		point.Lat,
		point.ID,
	)

	if err != nil {
		return nil, err
	}

	updatedPoint, err := r.GetByID(point.ID)

	if err != nil {
		return nil, err
	}

	if updatedPoint == nil {
		return nil, fmt.Errorf("Point with ID %d does not exists", point.ID)
	}

	return updatedPoint, nil
}

// GetAllByCity returns all points by city.
func (r *PointRepository) GetAllByCity(q queries.PointsQuery) ([]*domain.Point, error) {
	rows, err := r.store.db.Query(
		"SELECT p.id, p.name, p.description, p.lon, p.lat, pt.id, pt.name, c.id, c.name "+
			"FROM point p JOIN point_type pt ON p.type_id = pt.id JOIN city c ON p.city_id = c.id "+
			"WHERE c.id = ? "+
			"LIMIT ? OFFSET ?",
		q.City,
		q.Limit,
		q.Offset,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	points := make([]*domain.Point, 0)
	for rows.Next() {
		point := domain.NewEmptyPoint()
		err := rows.Scan(&point.ID, &point.Name, &point.Description, &point.Lon, &point.Lat, &point.Type.ID, &point.Type.Name, &point.City.ID, &point.City.Name)

		if err != nil {
			return nil, err
		}

		points = append(points, point)
	}

	return points, nil
}

// GetNearest returns nearest points by radius and coordinates.
func (r *PointRepository) GetNearest(q queries.NearestPointsQuery) ([]*domain.Point, error) {
	rows, err := r.store.db.Query(
		"SELECT p.id, p.name, p.description, p.lon, p.lat, pt.id, pt.name, c.id, c.name, get_distance(?, ?, p.lon, p.lat) as distance "+
			"FROM point p JOIN point_type pt ON p.type_id = pt.id JOIN city c ON p.city_id = c.id "+
			"HAVING distance <= ? "+
			"LIMIT ? OFFSET ?",
		q.Lon,
		q.Lat,
		q.Radius,
		q.Limit,
		q.Offset,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	points := make([]*domain.Point, 0)
	for rows.Next() {
		point := domain.NewEmptyPoint()
		err := rows.Scan(&point.ID, &point.Name, &point.Description, &point.Lon, &point.Lat, &point.Type.ID, &point.Type.Name, &point.City.ID, &point.City.Name, &point.Distance)

		if err != nil {
			return nil, err
		}

		points = append(points, point)
	}

	return points, nil
}

// GetByID returns the point by its identifier.
func (r *PointRepository) GetByID(id uint64) (*domain.Point, error) {
	point := domain.NewEmptyPoint()

	err := r.store.db.QueryRow(
		"SELECT p.id, p.name, p.description, p.lon, p.lat, pt.id, pt.name, c.id, c.name FROM point p JOIN point_type pt ON p.type_id = pt.id JOIN city c ON p.city_id = c.id WHERE p.id = ?;",
		id,
	).Scan(&point.ID, &point.Name, &point.Description, &point.Lon, &point.Lat, &point.Type.ID, &point.Type.Name, &point.City.ID, &point.City.Name)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return point, nil
}
