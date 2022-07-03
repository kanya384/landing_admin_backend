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
	fmt.Println(len(files))

	clientOptions := options.Client().ApplyURI(cfg.DSN)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Panic("error connection to mongo", err, nil)
	}

	cache := memcache.New(cfg)

	repository := mng.Setup(client.Database("public"))
	services := services.Setup(cfg, repository, logger, cache)

	err = services.Plans.ProcessPlans(context.Background(), files)
	if err != nil {
		fmt.Println(err)
		logger.Panic("error updating file", err, nil)
	}

	//update pdf files
	setting, err := services.Settings.GetByName(context.Background(), "remont_square_price")
	price := 0
	if err == nil {
		price = setting.Value
	}
	err = services.PdfGenerator.GeneratePdfs(context.Background(), price)
	if err == nil {
		fmt.Println("success")
	}

	ftp.CloseConnection()

}
