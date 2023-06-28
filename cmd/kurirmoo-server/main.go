package main

import (
	"kurirmoo"
	"kurirmoo/gen/restapi"
	"kurirmoo/gen/restapi/operations"
	"kurirmoo/internal/handlers"
	"kurirmoo/internal/rest"
	"log"
	"os"

	"github.com/casualjim/middlewares"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
	"github.com/justinas/alice"
	"github.com/spf13/viper"
)

var mainFlags = struct {
	AppConfig string `long:"config" description:"Main application configuration YAML path"`
}{}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	cfg := viper.New()
	cfg.SetConfigFile(".env")
	err := cfg.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed load config : %f", err)
	}

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalf("Error : %f", err)
	}

	api := operations.NewKurirmooServerAPI(swaggerSpec)
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "App Flags",
			LongDescription:  "",
			Options:          &mainFlags,
		},
	}

	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Kurirmoo Backend Service"
	parser.LongDescription = "Kurirmoo Backend Service"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	rt := kurirmoo.NewRuntime(cfg)

	h := handlers.NewHandler()

	rest.Authorization(rt, api)
	rest.Route(rt, api, h)

	logger := log.New(os.Stdout, "Log -> ", log.LstdFlags)

	api.Logger = func(s string, i ...interface{}) {
		logger.Println(s, i)
	}

	handler := alice.New(
		middlewares.NewRecoveryMW("kurirmoo", nil),
		middlewares.NewProfiler,
	).Then(api.Serve(nil))

	server.SetHandler(handler)
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
