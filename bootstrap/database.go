package bootstrap

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func connectDB(env *env) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		env.DBHost,
		env.DBUser,
		env.DBPass,
		env.DBName,
		env.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Can't connect to database: ", err)
	}

	return db
}

func closeDB(db *gorm.DB) {
	d, _ := db.DB()
	defer func() {
		if err := d.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
}

func autoMigration(db *gorm.DB) {
	err := db.AutoMigrate(
	//	TODO:: add domain's here
	)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("migration is done")
}
