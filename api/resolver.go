package api

import (
	"astro_pro/api/data"
	"astro_pro/api/data/mongodb"
	"astro_pro/api/models"
	rpc "astro_pro/rpc/pb"
	"context"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

var rep data.Repository
var rpcClient *RpcClient

func init() {
	rep = mongodb.NewRepository()
	rpcClient = NewRpcClient()
}

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r, rep, rpcClient}
}

type queryResolver struct {
	*Resolver
	repository data.Repository
	rpcClient  *RpcClient
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r, rpcClient}
}

type mutationResolver struct {
	*Resolver
	rpcClient *RpcClient
}

func (r *mutationResolver) SetDimming(ctx context.Context, mac string, value int) (string, error) {
	req := &rpc.Dimming{Mac: mac, Value: int32(value)}
	res, err := r.rpcClient.c.SetDimming(context.Background(), req)
	if err != nil {
		return "", err
	}
	return res.Message, nil
}

func (r *mutationResolver) SelectProfile(ctx context.Context, mac string, id int) (string, error) {
	req := &rpc.Profile{Mac: mac, Id: int32(id)}
	res, err := r.rpcClient.c.SelectProfile(context.Background(), req)
	if err != nil {
		return "", err
	}
	return res.Message, nil
}

func (r *mutationResolver) SetSimpleProfile(ctx context.Context, mac string, id int, profile models.SimpleProfileInput) (string, error) {
	req := &rpc.SimpleProfileRequest{
		Mac: mac,
		Id:  int32(id),
		Profile: &rpc.SimpleProfile{
			D1: int32(profile.D1),
			P1: int32(profile.P1),
			D2: int32(profile.D2),
			P2: int32(profile.P2),
		},
	}

	res, err := r.rpcClient.c.SetSimpleProfile(context.Background(), req)
	if err != nil {
		return "", err
	}

	return res.Message, nil
}
func (r *mutationResolver) SetComplexProfile(ctx context.Context, mac string, id int, profile models.ComplexProfileInput) (string, error) {
	req := &rpc.ComplexProfileRequest{
		Mac: mac,
		Id:  int32(id),
		Profile: &rpc.ComplexProfile{
			Pwm0:  int32(profile.Pwm0),
			Time1: profile.Time1,
			Pwm1:  int32(profile.Pwm1),
			Time2: profile.Time2,
			Pwm2:  int32(profile.Pwm2),
			Time3: profile.Time3,
			Pwm3:  int32(profile.Pwm3),
			Time4: profile.Time4,
			Pwm4:  int32(profile.Pwm4),
		},
	}

	res, err := r.rpcClient.c.SetComplexProfile(context.Background(), req)
	if err != nil {
		return "", err
	}

	return res.Message, nil
}

func (r *queryResolver) Cities(ctx context.Context) ([]*models.City, error) {
	return r.repository.Cities()
}

func (r *queryResolver) Segments(ctx context.Context, city *string) ([]*models.Segment, error) {
	return r.repository.Segments(city)
}

func (r *queryResolver) Controllers(ctx context.Context, city *string) ([]*models.Controller, error) {
	return r.repository.Controllers(city)
}

func (r *queryResolver) ControllersForSegment(ctx context.Context, city string, segment string) ([]*models.Controller, error) {
	return r.repository.ControllersForSegment(city, segment)
}

func (r *queryResolver) Controller(ctx context.Context, mac string) (*models.Controller, error) {
	return r.repository.Controller(mac)
}

func (r *queryResolver) Lamps(ctx context.Context, city *string) ([]*models.Lamp, error) {
	return r.repository.Lamps(city)
}

func (r *queryResolver) LampsForSegment(ctx context.Context, city string, segment string) ([]*models.Lamp, error) {
	return r.repository.LampsForSegment(city, segment)
}

func (r *queryResolver) Lamp(ctx context.Context, mac string) (*models.Lamp, error) {
	return r.repository.Lamp(mac)
}

func (r *queryResolver) LampNearByDistance(ctx context.Context, lat float64, lng float64, maxMeters int) (*models.Lamp, error) {
	return r.repository.LampNearByDistance(lat, lng, maxMeters)
}

func (r *queryResolver) LampsNearByCount(ctx context.Context, lat float64, lng float64, maxCount int) ([]*models.Lamp, error) {
	return r.repository.LampsNearByCount(lat, lng, maxCount)
}

func (r *queryResolver) RequestComplexProfile(ctx context.Context, mac string, id int) (string, error) {
	req := &rpc.Profile{Mac: mac, Id: int32(id)}
	res, err := r.rpcClient.c.RequestComplexProfile(context.Background(), req)
	if err != nil {
		return "", err
	}
	return res.Message, nil
}
