package rest

import (
	"kurirmoo"
	"kurirmoo/gen/models"
	"kurirmoo/gen/restapi/operations"
	"kurirmoo/internal/utils/jwt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Authorization(rt *kurirmoo.Runtime, api *operations.KurirmooServerAPI) {
	api.HasRoleAuth = func(token string, scopes []string) (*models.Principal, error) {
		// Load file .env
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}

		// Encode token
		signingKey, _ := jwt.NewJWTMaker(os.Getenv("JWT_SECRET"))
		payload, err := jwt.Maker.VerifyToken(signingKey, token)
		if err != nil {
			return nil, rt.SetError(http.StatusUnauthorized, err.Error())
		}

		// Cek role
		for _, scope := range scopes {
			if payload.Role == scope {
				return &models.Principal{
					Roles: []string{payload.Role},
				}, nil
			}
		}

		return nil, rt.SetError(http.StatusForbidden, "forbidden. only super admin and admin partner is allowed")
	}
}
