package database

import (
	"github.com/junyoung/fiber-angular/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

// Database instance
type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// Connect function
func Connect() {
	//p := Config("DB_PORT")
	// because our config function returns a string, we are parsing our      str to int here
	//port, err := strconv.ParseUint(p, 10, 32)
	//if err != nil {
	//	fmt.Println("Error parsing str to int")
	//}
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&model.User{})
	DB = Dbinstance{
		Db: db,
	}
}
