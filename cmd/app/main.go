package main

import (
	"fmt"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/generated"
	"landing_admin_backend/internal/generated/operations"
	"landing_admin_backend/pkg/logger"
	"log"
	"os"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
)

const (
	LOG_FILE = "./var/log/main.log"
)

func main() {
	cfg, err := config.InitConfig("APP")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}
	fileLog, err := os.OpenFile("./vars/logs/main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer fileLog.Close()

	logger, err := logger.NewLogger(cfg.ServiceName, cfg.LogLevel, LOG_FILE)
	swaggerSpec, err := loads.Embedded(generated.SwaggerJSON, generated.FlatSwaggerJSON)
	if err != nil {
		logger.Panic("error load swagger spec", err, nil)
	}

	api := operations.NewBackendServiceAPI(swaggerSpec)
	server := generated.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Landing Backend REST Api"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			logger.Panic("error adding swagger group", err, nil)
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
	server.Port = cfg.Port

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
