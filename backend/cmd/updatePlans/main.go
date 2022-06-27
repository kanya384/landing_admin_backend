package main

import (
	"context"
	"fmt"
	"landing_admin_backend/internal/config"
	mng "landing_admin_backend/internal/repository/mongo"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/internal/services/memcache"
	"landing_admin_backend/pkg/ftpfiles"
	"landing_admin_backend/pkg/helpers"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//mng "landing_admin_backend/internal/repository/mongo"
)

func main() {
	cfg, err := config.InitConfig("APP")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}

	file, err := os.OpenFile(cfg.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening log file config %s", err))
	}
	defer file.Close()

	logger, err := helpers.ConfigureLogger(cfg.ServiceName, cfg.LogLevel, file)
	if err != nil {
		panic(fmt.Sprintf("error initializing logger %s", err))
	}

	ftp, err := ftpfiles.NewSFtpFiles(cfg.FtpHost, cfg.FtpLogin, cfg.FtpPass, cfg.FtpPath)
	if err != nil {
		logger.Panic("error connectiong to ftp", err, nil)
	}

	files, err := ftp.GetFiles()
	if err != nil {
		logger.Panic("error getting files from ftp", err, nil)
	}
	if len(files) == 0 {
		os.Exit(0)
	}

	clientOptions := options.Client().ApplyURI(cfg.DSN)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Panic("error connection to mongo", err, nil)
	}

	cache := memcache.New()

	repository := mng.Setup(client.Database("public"))
	services := services.Setup(cfg, repository, logger, cache)

	err = services.Plans.ProcessPlans(context.Background(), files)
	if err != nil {
		logger.Panic("error updating file", err, nil)
	}

	ftp.CloseConnection()

}