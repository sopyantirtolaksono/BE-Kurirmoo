package handlers

import (
	"context"
	"kurirmoo"
	"kurirmoo/gen/models"
	"net/http"
	"regexp"
)

func isValidPhoneNumberFormat(phone string, rt *kurirmoo.Runtime) (bool, error) {
	match, err := regexp.MatchString(`^\d+$`, phone)
	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, "internal server error")
		return false, err
	}
	return match, nil
}

func isValidPhoneNumberLength(phone string) bool {
	return len(phone) == 12 || len(phone) == 13
}

func (h *handler) UpdateDriver(ctx context.Context, rt *kurirmoo.Runtime, ID int64, PhoneNumber string, PhoneNumber2 string) (message string, err error) {
	// Deklarasikan models drivers
	var driver models.Driver

	// Cek ID user/driver ada tidak ?
	err = rt.DB().WithContext(ctx).First(&driver, ID).Error
	if err != nil {
		err = rt.SetError(http.StatusNotFound, "user not found")
		return
	}

	// Cek phone number 2 ada tidak ?
	if PhoneNumber2 == "" {
		// Validasi format phone number
		if valid, er := isValidPhoneNumberFormat(PhoneNumber, rt); er != nil {
			err = rt.SetError(http.StatusInternalServerError, "internal server error")
			return
		} else if !valid {
			err = rt.SetError(http.StatusBadRequest, "invalid phone number format")
			return
		}
		// Validasi panjang phone number
		if !isValidPhoneNumberLength(PhoneNumber) {
			err = rt.SetError(http.StatusBadRequest, "invalid phone number length")
			return
		}

		driver.User.PhoneNumber = &PhoneNumber
	} else {
		// Validasi format phone number 1
		if valid, er := isValidPhoneNumberFormat(PhoneNumber, rt); er != nil {
			err = rt.SetError(http.StatusInternalServerError, "internal server error")
			return
		} else if !valid {
			err = rt.SetError(http.StatusBadRequest, "invalid phone number format")
			return
		}
		// Validasi panjang phone number 1
		if !isValidPhoneNumberLength(PhoneNumber) {
			err = rt.SetError(http.StatusBadRequest, "invalid phone number length")
			return
		}
		// Validasi format phone number 2
		if valid, er := isValidPhoneNumberFormat(PhoneNumber2, rt); er != nil {
			err = rt.SetError(http.StatusInternalServerError, "internal server error")
			return
		} else if !valid {
			err = rt.SetError(http.StatusBadRequest, "invalid phone number format")
			return
		}
		// Validasi panjang phone number 2
		if !isValidPhoneNumberLength(PhoneNumber2) {
			err = rt.SetError(http.StatusBadRequest, "invalid phone number length")
			return
		}

		result := PhoneNumber + "," + PhoneNumber2
		driver.User.PhoneNumber = &result
	}

	// Update phone number berdasarkan ID
	// pada tabel users
	err = rt.DB().WithContext(ctx).Model(&driver.User).Update("phone_number", &driver.User.PhoneNumber).Error
	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, "failed to update")
		return
	}
	// pada tabel drivers
	err = rt.DB().WithContext(ctx).Model(&driver).Update("phone_number", &driver.User.PhoneNumber).Error
	if err != nil {
		err = rt.SetError(http.StatusInternalServerError, "failed to update")
		return
	}

	// Return value jika sukses
	return "updated", nil
}
