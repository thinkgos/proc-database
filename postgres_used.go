//go:build postgres
// +build postgres

package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(dsn string) gorm.Dialector {
	return postgres.New(postgres.Config{DSN: dsn})
}
