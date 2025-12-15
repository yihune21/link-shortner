package main

func main()  {


	cfg := config{
		addr:":3333",
		db: dbConfig{
			dsn: "",
		},
	}
	app := appilication{
		config: cfg,
	}

	h := app.mount()
    app.run(h)


}