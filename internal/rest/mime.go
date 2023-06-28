package rest

import (
	"kurirmoo/gen/restapi/operations"

	"github.com/go-openapi/runtime"
)

func Mime(api *operations.KurirmooServerAPI) {
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
}
