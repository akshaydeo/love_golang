package dbWrapper

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

var db gorm.DB

func DB() *gorm.DB {
	return &db
}

func IntializeDatabase() error {
	switch os.Getenv("ENV") {
	case "dev":
		db, _ = gorm.Open("postgres", os.Getenv("DB_PATH"))
		break
	}
	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	db.DB()
	// Then you could invoke `*sql.DB`'s functions with it
	err := db.DB().Ping()
	if err != nil {
		return err
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// Disable table name's pluralization
	db.SingularTable(false)
	return nil
}
