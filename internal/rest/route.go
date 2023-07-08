package rest

import (
	"context"
	"kurirmoo"
	"kurirmoo/gen/models"
	"kurirmoo/gen/restapi/operations"
	"kurirmoo/gen/restapi/operations/add_city"
	"kurirmoo/gen/restapi/operations/cities"
	"kurirmoo/gen/restapi/operations/city_by_name"
	"kurirmoo/gen/restapi/operations/health"
	"kurirmoo/gen/restapi/operations/login"

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

	api.AddCityAddCityHandler = add_city.AddCityHandlerFunc(func(acp add_city.AddCityParams) middleware.Responder {
		err := apiHandler.AddCity(context.Background(), rt, *acp.Data.Name, *acp.Data.Code)
		
		if err != nil {
			errResponse := rt.GetError(err)
			return add_city.NewAddCityBadRequest().WithPayload(&models.Error{
				Code:    int64(errResponse.Code()),
				Message: errResponse.Error(),
			})
		}

		return add_city.NewAddCityCreated().WithPayload(&add_city.AddCityCreatedBody{
			Message: "Success to add city",
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
}
