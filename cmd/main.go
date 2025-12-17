package main

import (
	"log/slog"
	"os"

	"github.com/yihune21/link-shortner/internal/env"
	"github.com/yihune21/link-shortner/internal/utils"
)

func main()  {
    env.LoadEnv()
    port:= env.GetEnv("PORT")
	dbUrl := env.GetEnv("DB_URL")

	cfg := config{
		addr:port,
		db: dbConfig{
			dsn:dbUrl,
		},
	}
	logger :=slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dbConn,err := utils.ConnectDb(cfg.db.dsn)
	if err != nil {
		slog.Error("Failed to connect the db","error",err)
		os.Exit(1)
	}

    logger.Info("DB connected")

	app := appilication{
		config: cfg,
		db: dbConn,
	}
	 
    if err := app.run(app.mount()); err != nil{
		slog.Error("Failed to start the server","error",err)
		os.Exit(1)
	}


}