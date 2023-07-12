package handlers

import (
	"context"
	"kurirmoo"
	"kurirmoo/gen/models"
	"net/http"
)

func (h *handler) UpdateTripStatus(ctx context.Context, rt *kurirmoo.Runtime, ID int64) (message string, err error) {
	var driver models.Driver
	err = rt.DB().WithContext(ctx).Where("id = ?", ID).First(&driver).Error
	if err != nil {
		err = rt.SetError(http.StatusUnauthorized, "You don't have access")
		return "", err
	}

	if driver.TripStatus == "online" {
		driver.TripStatus = "offline"
	} else {
		driver.TripStatus = "online"
	}

	err = rt.DB().WithContext(ctx).Model(&driver).Update("trip_status", driver.TripStatus).Error
	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, "Internal server error")
		return "", err
	}

	return driver.TripStatus, nil
}
