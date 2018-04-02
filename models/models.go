package models

import (
	//"github.com/gilmarpalega/mlog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"time"
)

var db *gorm.DB

//var err error

type Flash struct {
	Type    string
	Message string
}

// Response contains the attributes found in an API response
type Response struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// Setup initializes the Conn object
// It also populates the Gophish Config object
func Setup() error {
	var err error
	db, err = gorm.Open("mysql", "root:@tcp(localhost:3306)/realiza?charset=utf8&parseTime=True&loc=America%2FSao_Paulo")
	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// Disable table name's pluralization
	db.SingularTable(true)
	db.LogMode(true)

	return err

}

func ForceOrder(order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

func GetDB() *gorm.DB {
	return db
}
