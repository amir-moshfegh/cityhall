package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env *env
	DB  *gorm.DB
}

func NewApplication() *Application {
	var app Application
	app.Env = newEnv()
	app.DB = connectDB(app.Env)

	return &app
}

func (app Application) CloseDB() {
	closeDB(app.DB)
}

func (app Application) AutoMigration() {
	autoMigration(app.DB)
}
