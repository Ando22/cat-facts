package database

import (
	"database/sql"
	"log"

	"github.com/Ando22/user-message/backend/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbHandler *sql.DB

// InitDatabase initializes the SQL database connection
func InitDatabase() error {
	connStr := config.DatabaseUsername + ":" + config.DatabasePassword + "@tcp(" + config.DatabaseHost + ":" + config.DatabasePort + ")/" + config.DatabaseName + "?parseTime=true"

	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return err
	}

	dbHandler = db.DB()

	log.Println("Connected to database")
	return nil
}

// GetDB returns the SQL database instance
func GetDB() *sql.DB {
	return dbHandler
}
