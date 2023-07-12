package handlers

import (
	"context"
	"kurirmoo"
	"kurirmoo/gen/models"
	"log"
	"net/http"
)

func (h *handler) UpdateTripStatus(ctx context.Context, rt *kurirmoo.Runtime, ID int64) (message string, err error) {
	var driver models.Driver
	err = rt.DB().WithContext(ctx).Where("id = ?", ID).First(&driver).Error
	if err != nil {
		err = rt.SetError(http.StatusNotFound, "Data not found")
		return "", err
	}

	log.Println("===>>>", driver.TripStatus)

	return driver.TripStatus, nil
}
