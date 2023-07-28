package handlers

import (
	"context"
	"errors"
	"kurirmoo"
	"kurirmoo/gen/models"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func IsUsed(id int64, arrDriverTruck []models.DriverTruck) bool {
	for _, value := range arrDriverTruck {
		if value.TruckID == uint64(id) {
			return true
		}
	}

	return false
}

func (h *handler) DeleteTruck(ctx context.Context, rt *kurirmoo.Runtime, ID int64) (message string, err error) {
	var driverTruck []models.DriverTruck
	var truck models.TruckList

	// Joins table driver_trucks and drivers where trip_status = online(active)
	err = rt.DB().WithContext(ctx).Table("driver_trucks").Select("driver_id", "truck_id").Joins("JOIN drivers ON driver_trucks.driver_id = drivers.id").Where("drivers.trip_status = ?", "online").Find(&driverTruck).Error
	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	// Check total active or online driver
	driverTruckLen := len(driverTruck)

	if driverTruckLen < 1 {
		// If truck ID not exists
		err = rt.DB().WithContext(ctx).First(&truck, ID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = rt.SetError(http.StatusNotFound, "not found")
				return
			}

			err = rt.SetError(http.StatusInternalServerError, err.Error())
			return
		}

		// Delete truck by ID
		// err = rt.DB().WithContext(ctx).Delete(&truck).Error
		currentTime := time.Now().Unix()
		err = rt.DB().WithContext(ctx).Model(&truck).Update("deleted_at", currentTime).Error
		if err != nil {
			err = rt.SetError(http.StatusInternalServerError, err.Error())
			return
		}

		// Return message success to delete truck
		return "deleted", nil
	} else {
		// If truck is used
		if IsUsed(ID, driverTruck) {
			return "can't delete! truck is used", nil
		}

		// If truck ID not exists
		err = rt.DB().WithContext(ctx).First(&truck, ID).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = rt.SetError(http.StatusNotFound, "not found")
				return
			}

			err = rt.SetError(http.StatusInternalServerError, err.Error())
			return
		}

		// Delete truck by ID
		// err = rt.DB().WithContext(ctx).Delete(&truck).Error
		currentTime := time.Now().Unix()
		err = rt.DB().WithContext(ctx).Model(&truck).Update("deleted_at", currentTime).Error
		if err != nil {
			err = rt.SetError(http.StatusInternalServerError, err.Error())
			return
		}

		// Return message success to delete truck
		return "deleted", nil
	}
}
