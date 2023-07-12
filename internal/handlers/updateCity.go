package handlers

import (
	"context"
	"kurirmoo"
	"kurirmoo/gen/models"
	"net/http"
)

func (h *handler) UpdateCity(ctx context.Context, rt *kurirmoo.Runtime, ID int64, name string, code string) (cityID int64, cityName string, cityCode string, err error) {

	var existingCity models.City
	err = rt.DB().WithContext(ctx).Where("id = ?", ID).First(&existingCity).Error
	if err != nil {
		err = rt.SetError(http.StatusNotFound, "Data not found")
		return 0, "", "", err
	}

	existingCity.CityName = name
	existingCity.CityCode = code

	err = rt.DB().WithContext(ctx).Save(&existingCity).Error
	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, "Internal server error")
		return 0, "", "", err
	}

	return int64(*existingCity.ID), existingCity.CityName, existingCity.CityCode, nil
}
