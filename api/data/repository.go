package data

import (
	"astro_pro/api/models"
)

type Repository interface {
	Cities() ([]*models.City, error)
	Segments(city *string) ([]*models.Segment, error)
	ControllersForSegment(city string, segment string) ([]*models.Controller, error)
	Controllers(city *string) ([]*models.Controller, error)
	Controller(mac string) (*models.Controller, error)
	Lamps(city *string) ([]*models.Lamp, error)
	LampsForSegment(city string, segment string) ([]*models.Lamp, error)
	Lamp(mac string) (*models.Lamp, error)
	LampNearByDistance(lat float64, lon float64, maxMeters int) (*models.Lamp, error)
	LampsNearByCount(lat float64, lng float64, maxCount int) ([]*models.Lamp, error)
}
