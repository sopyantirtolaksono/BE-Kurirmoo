package handlers

import (
	"context"
	"kurirmoo"
	"kurirmoo/gen/models"
	"net/http"
)

type OriginAndDestination struct {
	OriginCity      string
	DestinationCity string
}

type CityPasseds struct {
	City string
}

func (h *handler) GetRouteAndCityPasseds(ctx context.Context, rt *kurirmoo.Runtime, ID int64) (origin_city string, destination_city string, city_passeds []string, err error) {

	var cities []string
	var route models.Route
	var cityPass models.CityPassed
	var originAndDestination OriginAndDestination
	var cityPasseds []CityPasseds

	err = rt.DB().WithContext(ctx).Model(&route).Where("id = ?", ID).Select("origin_city", "destination_city").Scan(&originAndDestination).Error

	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, "Internal server error")
		return
	}

	err = rt.DB().WithContext(ctx).Model(&cityPass).Where("route_id = ?", ID).Select("city").Scan(&cityPasseds).Error

	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, "Internal server error")
		return
	}

	for _, cityPassed := range cityPasseds {
		cities = append(cities, cityPassed.City)
	}

	return originAndDestination.OriginCity, originAndDestination.DestinationCity, cities, nil
}
