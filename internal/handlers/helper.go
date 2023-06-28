package handlers

import (
	"context"
	"kurirmoo"
	"math"
	"net/http"
	"time"

	"kurirmoo/gen/models"

	"github.com/go-openapi/runtime"
	http_transport "github.com/go-openapi/runtime/client"
)

func newCustomTransport() *customTransport {
	return &customTransport{
		originalTransport: http.DefaultTransport,
	}
}

type customTransport struct {
	originalTransport http.RoundTripper
}

func (c *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := c.originalTransport.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getContextTimeout(parent context.Context, rt *kurirmoo.Runtime) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, time.Second*time.Duration(rt.Config().GetInt64("OTP_WHATSAPP_RTO")))
}

func getInternalOtpUrl(rt *kurirmoo.Runtime, url string) string {
	return url
}

func getInternalOtpAuth(secret string) runtime.ClientAuthInfoWriter {
	return http_transport.APIKeyAuth("x-api-key", "header", secret)
}

// getPaginationData is helper to get pagination data
func getPaginationData(page int, limit float64, totalRow float64) *models.Metadata {
	totalPage := math.Ceil(totalRow / limit)

	return &models.Metadata{
		Page:      int64(page),
		TotalRow:  int64(totalRow),
		TotalPage: int64(totalPage),
		PerPage:   int64(limit),
	}
}
