package main

func main() {
	/*cfg, err := config.InitConfig("APP")
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

	clientOptions := options.Client().ApplyURI(cfg.DSN)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Panic("error connection to mongo", err, nil)
	}

	cache := memcache.New()

	repository := mng.Setup(client.Database("public"))
	services := services.Setup(cfg, repository, logger, cache)

	err = services.PdfGenerator.GeneratePdfs(context.Background())
	if err == nil {
		fmt.Println("success")
	}*/

}
