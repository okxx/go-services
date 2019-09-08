package mysql

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

var DB *gorm.DB

type database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}
var Driver = &database{}

func Setup() {
	var err error
	DB, err = gorm.Open(Driver.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		Driver.User,
		Driver.Password,
		Driver.Host,
		Driver.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return Driver.TablePrefix + defaultTableName
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func ClonseDB() {
	defer DB.Close()
}
