package handlers

import (
	"context"
	"kurirmoo"
	"kurirmoo/gen/models"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

func (h *handler) CityByName(ctx context.Context, rt *kurirmoo.Runtime, name string) ( city_name string, city_code string, acronim string, err error) {
	cities := &models.City{}

	result := rt.DB().WithContext(ctx).Model(cities).Where("city_name = ?", name).First(cities)
	err = result.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = rt.SetError(http.StatusNotFound, "No record found in the database")
		}
	}


	city_name = *&cities.CityName
	city_code = *&cities.CityCode

	
	words := strings.Split(city_name, " ")
	for _, word := range words {
    acronim = acronim + string([]rune(word)[0])
	}

	return  city_name, city_code, acronim, err
}
