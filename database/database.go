package database

import (
	"fmt"
	"log"
	"os"

	"github.com/DjadiHadia/restapiwithgo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Agency{})
	db.AutoMigrate(&models.Car{})
	db.AutoMigrate(&models.Person{})
	db.AutoMigrate(&models.Client{})
	db.AutoMigrate(&models.Reservation{})
	db.AutoMigrate(&models.User{})

	DB = Dbinstance{
		Db: db,
	}
	/*
		// Query to retrieve table names from information_schema.tables
		var tableNames []string
		if err := db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Pluck("table_name", &tableNames).Error; err != nil {
			log.Fatal("Failed to retrieve table names:", err)
		}

		// Output the table names
		log.Println("Tables in the database:")
		for _, tableName := range tableNames {
			log.Println(tableName)
		}*/

}

// test queries
// TestDB executes test queries
func TestDB(db *gorm.DB) {
	// Create an instance of the GORM migrator
	migrator := db.Migrator()

	// Retrieve table information using GORM's migrator
	tables := migrator.CurrentDatabase()

	log.Println("Tables in the database:")
	for _, table := range tables {
		log.Println(table)
	}
}
