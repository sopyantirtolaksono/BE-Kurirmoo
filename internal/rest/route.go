package rest

import (
	"context"
	"kurirmoo"
	"kurirmoo/gen/models"
	"kurirmoo/gen/restapi/operations"
	"kurirmoo/gen/restapi/operations/cities"
	"kurirmoo/gen/restapi/operations/city_by_name"
	"kurirmoo/gen/restapi/operations/health"
	"kurirmoo/gen/restapi/operations/login"
	"kurirmoo/gen/restapi/operations/trucks"

	"kurirmoo/internal/handlers"

	"github.com/go-openapi/runtime/middleware"
)

func Route(rt *kurirmoo.Runtime, api *operations.KurirmooServerAPI, apiHandler handlers.Handler) {
	api.LoginAuthHandler = login.AuthHandlerFunc(func(params login.AuthParams) middleware.Responder {
		_, token, expiredAt, err := apiHandler.Login(context.Background(), rt, *params.Data.Email, *params.Data.Password)
		if err != nil {
			errResponse := rt.GetError(err)
			return login.NewAuthBadRequest().WithPayload(&models.Error{
				Code:    int64(errResponse.Code()),
				Message: errResponse.Error(),
			})
		}

		return login.NewAuthOK().WithPayload(&login.AuthOKBody{
			Message:   "Success to Login",
			Token:     token,
			ExpiredAt: expiredAt,
		})
	})

	api.HealthHealthHandler = health.HealthHandlerFunc(func(params health.HealthParams) middleware.Responder {
		HealthCheck := apiHandler.Health()

		return health.NewHealthOK().WithPayload(&models.Success{Message: HealthCheck})
	})

	api.CitiesGetAllCitiesHandler = cities.GetAllCitiesHandlerFunc(func(params cities.GetAllCitiesParams) middleware.Responder {
		var cityList []*cities.GetAllCitiesOKBodyItems0

		cityList, err := apiHandler.Cities(context.Background(), rt, params)

		if err != nil {
			errResponse := rt.GetError(err)
			return city_by_name.NewGetCityByNameBadRequest().WithPayload(&models.Error{
				Code:    int64(errResponse.Code()),
				Message: errResponse.Error(),
			})
		}

		return cities.NewGetAllCitiesOK().WithPayload(cityList)
	})

	api.CityByNameGetCityByNameHandler = city_by_name.GetCityByNameHandlerFunc(func(params city_by_name.GetCityByNameParams) middleware.Responder {
		city_name, city_code, acronim, err := apiHandler.CityByName(context.Background(), rt, *&params.CityName)

		if err != nil {
			errResponse := rt.GetError(err)
			return city_by_name.NewGetCityByNameBadRequest().WithPayload(&models.Error{
				Code:    int64(errResponse.Code()),
				Message: errResponse.Error(),
			})
		}

		return city_by_name.NewGetCityByNameOK().WithPayload(&city_by_name.GetCityByNameOKBody{
			CityCode: city_code,
			CityName: city_name,
			Acronim:  acronim,
		})
	})

	api.TrucksAddTruckHandler = trucks.AddTruckHandlerFunc(func(params trucks.AddTruckParams) middleware.Responder {
		message, err := apiHandler.AddTruck(context.Background(), rt, *params.Data.TruckKind, *params.Data.Brand, *params.Data.TruckType, *params.Data.Length, *params.Data.Width, *params.Data.Height, *params.Data.Capacity)

		if err != nil {
			errResponse := rt.GetError(err)
			return trucks.NewAddTruckBadRequest().WithPayload(&models.Error{
				Code:    int64(errResponse.Code()),
				Message: errResponse.Error(),
			})
		}

		return trucks.NewAddTruckCreated().WithPayload(&models.Success{
			Message: message,
		})
	})
}
