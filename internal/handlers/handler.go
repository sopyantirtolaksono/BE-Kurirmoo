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
		AddTruckHandler
		DetailDataMultiplierHandler
		DeleteTruckHandler
	}

	LoginHandler interface {
		Login(ctx context.Context, rt *kurirmoo.Runtime, email, password string) (userId uint64, token, expiredAt string, err error)
	}

	CitiesHandler interface {
		Cities(ctx context.Context, rt *kurirmoo.Runtime, params cities.GetAllCitiesParams) (cityList []*cities.GetAllCitiesOKBodyItems0, err error)
	}

	CityByNameHandler interface {
		CityByName(ctx context.Context, rt *kurirmoo.Runtime, name string) (city_name string, city_code string, acronim string, err error)
	}

	HealthHandler interface {
		Health() string
	}

	AddTruckHandler interface {
		AddTruck(ctx context.Context, rt *kurirmoo.Runtime, truck_kind string, brand string, truck_type string, length int64, width int64, height int64, capacity int64) (message string, err error)
	}

	DetailDataMultiplierHandler interface {
		DetailDataMultiplier(ctx context.Context, rt *kurirmoo.Runtime, ID string) (route string, city_passed int64, lane string, distance int64, truck_kind string, brand string, truck_type string, volume int64, capacity int64, max_price int64, min_price int64, price_per_km int64, id int64, err error)
	}

	DeleteTruckHandler interface {
		DeleteTruck(ctx context.Context, rt *kurirmoo.Runtime, ID int64) (message string, err error)
	}
)

func NewHandler() Handler {
	return &handler{}
}

type handler struct{}
