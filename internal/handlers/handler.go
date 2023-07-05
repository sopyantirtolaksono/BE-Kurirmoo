package handlers

import (
	"context"
	"kurirmoo"

	"kurirmoo/gen/restapi/operations/cities"
)

type (
	Handler interface {
		LoginHandler
		CitiesHandler
		CityByNameHandler
		HealthHandler
		AddCityHandler
	}

	LoginHandler interface {
		Login(ctx context.Context, rt *kurirmoo.Runtime, email, password string) (userId uint64, token, expiredAt string, err error)
	}

	CitiesHandler interface {
		Cities(ctx context.Context, rt *kurirmoo.Runtime, params cities.GetAllCitiesParams) (cityList []*cities.GetAllCitiesOKBodyItems0, err error)
	}

	AddCityHandler interface {
		AddCity(ctx context.Context, rt *kurirmoo.Runtime, name, code string) (err error)
	}

	CityByNameHandler interface {
		CityByName(ctx context.Context, rt *kurirmoo.Runtime, name string) (city_name string, city_code string, acronim string, err error)
	}

	HealthHandler interface {
		Health() string
	}
)

func NewHandler() Handler {
	return &handler{}
}

type handler struct{}
