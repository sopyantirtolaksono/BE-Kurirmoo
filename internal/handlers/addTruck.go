package handlers

import (
	"context"
	"kurirmoo"
	"kurirmoo/gen/models"
)

func (h *handler) AddTruck(ctx context.Context, rt *kurirmoo.Runtime, truck_kind string, brand string, truck_type string, length int64, width int64, height int64, capacity int64) (message string, err error) {

	truck := &models.TruckList{}

	truck.TruckKind = &truck_kind
	truck.Brand = &brand
	truck.TruckType = &truck_type
	truck.Length = &length
	truck.Width = &width
	truck.Height = &height
	truck.Capacity = &capacity

	result := rt.DB().WithContext(ctx).Create(&truck)
	if result.Error != nil {
		return "Failed to add truck", result.Error
	}

	return "Success added", nil
}
