package mongodb

import (
	"astro_pro/api/models"
	mg "astro_pro/data/mongodb"
)

type Repository struct {
	repository *mg.Repository
}

func NewRepository() *Repository {
	return &Repository{repository: mg.NewRepository()}
}

func (r *Repository) Cities() ([]*models.City, error) {
	cs, err := r.repository.Cities()
	if err != nil {
		return nil, err
	}

	var cities []*models.City

	for _, c := range cs {
		cities = append(cities, cityToDomain(c))
	}

	return cities, nil
}

func (r *Repository) Segments(city *string) ([]*models.Segment, error) {
	sgs, err := r.repository.Segments(city)
	if err != nil {
		return nil, err
	}

	var segments []*models.Segment

	for _, s := range sgs {
		segments = append(segments, segmentToDomain(s))
	}

	return segments, nil
}

func (r *Repository) ControllersForSegment(city string, segment string) ([]*models.Controller, error) {
	ctrls, err := r.repository.ControllersForSegment(city, segment)
	if err != nil {
		return nil, err
	}

	var controllers []*models.Controller

	for _, c := range ctrls {
		controllers = append(controllers, controllerToDomain(c))
	}

	return controllers, nil
}

func (r *Repository) Controllers(city *string) ([]*models.Controller, error) {
	ctrls, err := r.repository.Controllers(city)
	if err != nil {
		return nil, err
	}

	var controllers []*models.Controller

	for _, c := range ctrls {
		controllers = append(controllers, controllerToDomain(c))
	}

	return controllers, nil
}

func (r *Repository) Controller(mac string) (*models.Controller, error) {
	c, err := r.repository.Controller(mac)
	if err != nil {
		return nil, err
	}

	return controllerToDomain(c), nil
}

func (r *Repository) LampsForSegment(city string, segment string) ([]*models.Lamp, error) {
	lps, err := r.repository.LampsForSegment(city, segment)
	if err != nil {
		return nil, err
	}

	var lamps []*models.Lamp

	for _, l := range lps {
		lamps = append(lamps, lampToDomain(l))
	}

	return lamps, nil
}

func (r *Repository) LampNearByDistance(lat float64, lon float64, maxMeters int) (*models.Lamp, error) {
	l, err := r.repository.LampNearByDistance(lat, lon, maxMeters)
	if err != nil {
		return nil, err
	}

	if l != nil {
		return lampToDomain(l), nil
	}

	return nil, nil
}

func (r *Repository) LampsNearByCount(lat float64, lng float64, maxCount int) ([]*models.Lamp, error) {
	lps, err := r.repository.LampsNearByCount(lat, lng, maxCount)
	if err != nil {
		return nil, err
	}

	var lamps []*models.Lamp

	for _, l := range lps {
		lamps = append(lamps, lampToDomain(l))
	}

	return lamps, nil
}

func (r *Repository) Lamps(city *string) ([]*models.Lamp, error) {
	lps, err := r.repository.Lamps(city)
	if err != nil {
		return nil, err
	}

	var lamps []*models.Lamp

	for _, l := range lps {
		lamps = append(lamps, lampToDomain(l))
	}

	return lamps, nil
}

func (r *Repository) Lamp(mac string) (*models.Lamp, error) {
	l, err := r.repository.Lamp(mac)
	if err != nil {
		return nil, err
	}

	if l != nil {
		return lampToDomain(l), nil
	}

	return nil, nil
}
