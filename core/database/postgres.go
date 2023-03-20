package database

import (
	"fmt"

	"github.com/dduafa/go-server/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func NewDBInstance(cfg *core.Config) (*gorm.DB, error) {
	var (
		dbInstance *gorm.DB
		err        error
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.PG_HOST, cfg.PG_USER, cfg.PG_PASS, cfg.PG_DB_NAME, cfg.PG_PORT, cfg.PG_SSLMODE)

	dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	dbInstance = dbInstance.Debug()

	fmt.Println("? Connected Successfully to the Database")
	return dbInstance, nil
}
