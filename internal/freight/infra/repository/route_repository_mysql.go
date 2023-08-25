package repository

import (
	"database/sql"

	"github.com/wsldev/gofreight/internal/freight/entity"
)

type RouteRepositoryMysql struct {
	db *sql.DB
}

func NewRouteRepositoryMysql(db *sql.DB) *RouteRepositoryMysql {
	return &RouteRepositoryMysql{
		db: db,
	}
}

func (r *RouteRepositoryMysql) Create(route *entity.Route) error {
	stmt := "INSERT INTO routes (id, name, distance, status, freight_price) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.Exec(stmt, route.ID, route.Name, route.Distance, route.Status, route.FreightPrice)

	if err != nil {
		return err
	}
	return nil
}

func (r *RouteRepositoryMysql) FindById(id string) (*entity.Route, error) {
	stmt := "SELECT id, name, distance, status, freight_price started_at, finished_at FROM routes WHERE id= ?"
	row := r.db.QueryRow(stmt, id)

	var startedAt, finishedAt sql.NullTime
	var route entity.Route

	err := row.Scan(
		&route.ID,
		&route.Name,
		&route.Distance,
		&route.Status,
		&route.FreightPrice,
		&startedAt,
		&finishedAt,
	)
	if err != nil {
		return nil, err
	}

	if startedAt.Valid {
		route.StartedAt = startedAt.Time
	}

	if finishedAt.Valid {
		route.FinishedAt = finishedAt.Time
	}

	return &route, nil
}

func (r *RouteRepositoryMysql) Update(route *entity.Route) error {
	startedAt := route.StartedAt.Format("2006-01-02T15:04")
	finishedAt := route.FinishedAt.Format("2006-01-02T15:04")
	stmt := "UPDATE routes SET status = ?, fright_price = ?, started_at = ?, finished_at = ? WHERE id = ?"
	_, err := r.db.Exec(stmt, route.Status, route.FreightPrice, startedAt, finishedAt, route.ID)
	if err != nil {
		return err
	}
	return nil
}
