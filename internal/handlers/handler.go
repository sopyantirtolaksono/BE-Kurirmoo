package handlers

import (
	"context"
	"kurirmoo"

	"kurirmoo/gen/restapi/operations/cities"
)

type (
	Handler interface {
		LoginHandler
<<<<<<< HEAD
		CitiesHandler
		CityByNameHandler
=======
		HealthHandler
>>>>>>> 5f48244841e9d3d5fd8c9c1e5bb243355dd95ce5
	}

	LoginHandler interface {
		Login(ctx context.Context, rt *kurirmoo.Runtime, email, password string) (userId uint64, token, expiredAt string, err error)
	}
<<<<<<< HEAD

	CitiesHandler interface {
		Cities(ctx context.Context, rt *kurirmoo.Runtime, params cities.GetAllCitiesParams) (cityList []*cities.GetAllCitiesOKBodyItems0, err error)
	}

	CityByNameHandler interface {
		CityByName(ctx context.Context, rt *kurirmoo.Runtime, name string) ( city_name string, city_code string, acronim string, err error)
=======
	
	HealthHandler interface {
		Health() (string)
>>>>>>> 5f48244841e9d3d5fd8c9c1e5bb243355dd95ce5
	}
)

func NewHandler() Handler {
	return &handler{}
}

type handler struct{}
