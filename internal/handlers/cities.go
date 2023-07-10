package handlers

import (
	"context"
	"kurirmoo"
	"time"

	"kurirmoo/gen/restapi/operations/cities"
	"net/http"

	"gorm.io/gorm"
)

func (h *handler) Cities(ctx context.Context, rt *kurirmoo.Runtime, params cities.GetAllCitiesParams) ([]*cities.GetAllCitiesOKBodyItems0, error) {

	if params.Limit == nil {
		*params.Limit = 10
	}

	var cityList []*cities.GetAllCitiesOKBodyItems0

	if params.CityName == nil {
		rt.Db.Raw("SELECT * FROM cities ORDER BY id ASC limit ?", params.Limit).Scan(&cityList)
	}

	if params.CityName != nil {
		rt.Db.Raw("SELECT * FROM cities WHERE city_name = ? limit ?", params.CityName, params.Limit).Scan(&cityList)
	}

	err := rt.Db.Raw("SELECT * FROM cities").Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = rt.SetError(http.StatusNotFound, "No record found in the database")
		}
	}

	return  cityList, err
}

func (h *handler) AddCity(ctx context.Context, rt *kurirmoo.Runtime, name, code string) (err error){

	if name == "" {
		err = rt.SetError(http.StatusBadRequest, "City's name is required")
		return err
	}

	if code == "" {
		err = rt.SetError(http.StatusBadRequest, "City's code is required")
		return err
	}

	createdAt := time.Now().Unix()

	result := rt.Db.Exec("INSERT INTO cities (city_name, city_code, created_at) VALUES (?, ?, ?)", name, code, createdAt)

	if result.Error != nil {
		err = rt.SetError(http.StatusNotFound, "Something went wrong. Adding data failed.")
		return err
	}


	return nil
}