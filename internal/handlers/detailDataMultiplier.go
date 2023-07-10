package handlers

import (
	"context"
	"kurirmoo"
	"net/http"
	"strconv"
	"strings"
)

type DetailMultiplier struct {
	ID              uint64
	RouteID         float64
	MinPrice        float64
	MaxPrice        float64
	TruckKind       string
	Brand           string
	TruckType       string
	Length          int64
	Width           int64
	Height          int64
	Capacity        int64
	DestinationCity string
	OriginCity      string
	Lane            string
	Distance        int64
}

func (h *handler) DetailDataMultiplier(ctx context.Context, rt *kurirmoo.Runtime, ID string) (route string, city_passed int64, lane string, distance int64, truck_kind string, brand string, truck_type string, volume int64, capacity int64, max_price int64, min_price int64, price_per_km int64, id int64, err error) {

	var detailMultiplier DetailMultiplier
	var cityPassed int64

	result := rt.DB().WithContext(ctx).Table("data_multipliers").Select("routes.destination_city, routes.origin_city, routes.lane, routes.distance, truck_lists.truck_kind, truck_lists.brand, truck_lists.truck_type, truck_lists.length, truck_lists.width, truck_lists.height, truck_lists.capacity, data_multipliers.min_price, data_multipliers.max_price, data_multipliers.route_id, data_multipliers.id").Joins("JOIN routes ON routes.id = data_multipliers.route_id").Joins("JOIN truck_lists ON truck_lists.id = data_multipliers.truck_id").Where("data_multipliers.id = ?", ID).Scan(&detailMultiplier)

	row := result.RowsAffected
	err = result.Error
	if row < 1 && err == nil {
		err = rt.SetError(http.StatusNotFound, "Data not found")
		return "", 0, "", 0, "", "", "", 0, 0, 0, 0, 0, 0, err
	} else if err != nil {
		err = rt.SetError(http.StatusBadRequest, "Bad request")
		return "", 0, "", 0, "", "", "", 0, 0, 0, 0, 0, 0, err
	}

	multiplierId := detailMultiplier.ID
	routeId := strconv.FormatFloat(detailMultiplier.RouteID, 'f', -1, 64)
	minPrice := int64(detailMultiplier.MinPrice)
	maxPrice := int64(detailMultiplier.MaxPrice)
	truckKind := detailMultiplier.TruckKind
	isBrand := detailMultiplier.Brand
	truckType := detailMultiplier.TruckType
	length := detailMultiplier.Length
	width := detailMultiplier.Width
	height := detailMultiplier.Height
	isCapacity := detailMultiplier.Capacity
	destinationCity := detailMultiplier.DestinationCity
	originCity := detailMultiplier.OriginCity
	isLane := detailMultiplier.Lane
	isDistance := detailMultiplier.Distance

	isRoute := destinationCity + " - " + originCity
	isVolume := length * width * height
	pricePerKm := maxPrice / isDistance

	rt.DB().WithContext(ctx).Table("city_passeds").Where("route_id = ?", routeId).Count(&cityPassed)

	return strings.ToUpper(isRoute), cityPassed, isLane, isDistance, truckKind, isBrand, truckType, isVolume, isCapacity, maxPrice, minPrice, pricePerKm, int64(multiplierId), nil
}
