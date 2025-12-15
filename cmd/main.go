package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/yihune21/link-shortner/internal/env"
	"github.com/yihune21/link-shortner/internal/utils"
)

func main()  {
    env.LoadEnv()
    port , dbUrl := env.GetEnv()
	cfg := config{
		addr:port,
		db: dbConfig{
			dsn:dbUrl,
		},
	}
	app := appilication{
		config: cfg,
	}
	 
	_,err := utils.ConnectDb(app.config.db.dsn)
	if err != nil {
	   	log.Fatalf ("Failed to connect the db %v",err)
		os.Exit(1)
	}
	println("Connected")
	logger :=slog.New(slog.NewJSONHandler(os.Stdout, nil))
    slog.SetDefault(logger)
    if err := app.run(app.mount()); err != nil{
		log.Fatalf ("Failed to start the server %v",err)
		os.Exit(1)
	}


}