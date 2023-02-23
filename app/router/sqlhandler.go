package router

import (
	"fmt"

	"github.com/SkyIsSkyBlue/ScheBlue-server/conf"
	"github.com/SkyIsSkyBlue/ScheBlue-server/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	DB *gorm.DB
}

func NewDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.SqlUsername(),
		conf.SqlPassword(),
		conf.SqlHostname(),
		conf.SqlPort(),
		conf.SqlDatabase(),
	)
	db, err = gorm.Open(
		mysql.New(mysql.Config{
			DSN: dsn,
		}),
		&gorm.Config{},
	)
	return
}

func migration(db *gorm.DB) error {
	return db.AutoMigrate(model.AllTables()...)
}

func initPing(db *gorm.DB) {
	db.Create(
		&model.SqlPing{
			PingId:    "sqlPing",
			PongValue: "sqlPong",
		})
}

func NewSqlHandler() (SqlHandler, error) {
	var sqlh SqlHandler
	var err error

	// init DB
	sqlh.DB, err = NewDB()
	if err != nil {
		return SqlHandler{}, err
	}

	// new DB
	err = migration(sqlh.DB)
	if err != nil {
		return SqlHandler{}, err
	}

	initPing(sqlh.DB)

	return sqlh, nil
}
