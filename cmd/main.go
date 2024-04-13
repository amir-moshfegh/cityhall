package main

import (
	"github.com/amir-moshfegh/cityhall/api/router"
	"github.com/amir-moshfegh/cityhall/bootstrap"
	"time"
)

func main() {
	app := bootstrap.NewApplication()
	r := router.New()

	if app.Env.AppEnv == "development" {
		app.DB.Debug()
	}

	timeOut := time.Duration(app.Env.ContextTimeout) * time.Second

	v1 := r.Group("/v1")
	router.Setup(app, v1, timeOut)

	r.Logger.Fatal(r.Start(app.Env.ServerAddress))
}
