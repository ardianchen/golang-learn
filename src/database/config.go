package database

import (
	"fmt"
	"os"

	mdl "../models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

// var validate *validator.Validate

// type Model struct {
// 	ID        uint       `gorm:"primary_key" json:"id,omitempty"`
// 	CreatedAt time.Time  `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
// 	UpdatedAt time.Time  `gorm:"not null" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
// 	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
// }

func init() {
	// check load env
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
		panic("check env file")
	}
	// config
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	db, err := gorm.Open("mysql", username+":"+password+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Asia%2FKolkata")

	if err != nil {
		fmt.Print(err)
		panic("failed to connect to database")
	}
	// db = conn

	//Printing query
	db.LogMode(true)

	//Automatically create migration as per model
	db.AutoMigrate(
		&mdl.User{},
	)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDb() {
	db.Close()
}
