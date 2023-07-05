package handlers

import (
	"context"
	"kurirmoo"
	"net/http"
)


func (h *handler) AddCity(ctx context.Context, rt *kurirmoo.Runtime, name, code string) (err error){

	if name == "" {
		err = rt.SetError(http.StatusNotFound, "City's name is required")
		return err
	}

	if code == "" {
		err = rt.SetError(http.StatusNotFound, "City's code is required")
		return err
	}

	result := rt.Db.Exec("INSERT INTO cities (city_name, city_code) VALUES (?, ?)", name, code)

	if result.Error != nil {
		err = rt.SetError(http.StatusNotFound, "Something went wrong. Adding data failed.")
		return err
	}


	return nil
}