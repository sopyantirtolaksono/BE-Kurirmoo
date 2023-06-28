package handlers

import (
	"context"
	"kurirmoo"
	"kurirmoo/gen/models"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func (h *handler) Login(ctx context.Context, rt *kurirmoo.Runtime, email, password string) (userId uint64, token, expiredAt string, err error) {
	user := &models.User{}

	err = rt.DB().WithContext(ctx).Model(user).Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = rt.SetError(http.StatusNotFound, "Email not registered yet")
		}
	}

	// TODO Check password is match or not

	// TODO Generate JWT Token
	generatedToken := "azajksnddajsdkansgkjaksjkdjaskdjaskdjaskdjkaksjdkasjda"

	// TODO Get expired at from JWT Token
	expiredTime := time.Now().Add(time.Hour * time.Duration(24))

	id := uint64(0)
	if user.ID != nil {
		id = *user.ID
	}

	return id, generatedToken, expiredTime.Format(time.RFC3339), err
}
