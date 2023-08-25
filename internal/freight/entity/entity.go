package entity

import "time"

type RouteRepository interface {
	Create(route *Route) error
	FindByID(id string) (*Route, error)
	Update(route *Route) error
}

type Route struct {
	ID           string
	Name         string
	Distance     float64
	Status       string
	FreightPrice float64
	StartedAt    time.Time
	FinishedAt   time.Time
}

type FreightInterface interface {
	Calculate(route *Route)
}

func NewRoute(id, name string, distance float64) *Route {
	return &Route{
		ID:       id,
		Name:     name,
		Distance: distance,
		Status:   "pending",
	}
}

func (r *Route) Start(startedAt time.Time) {
	r.Status = "started"
	r.StartedAt = startedAt
}

func (r *Route) Finish(finishedAt time.Time) {
	r.Status = "finished"
	r.FinishedAt = finishedAt
}

type Freight struct {
	PricePerKm float64
}

func NewFreight(pricePerKm float64) *Freight {
	return &Freight{
		PricePerKm: pricePerKm,
	}
}

func (f *Freight) Calculate(route *Route) {
	route.FreightPrice = route.Distance * f.PricePerKm
}
