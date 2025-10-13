//go:build !postgres
// +build !postgres

package database

import (
	"gorm.io/gorm"
)

func NewPostgres(dsn string) gorm.Dialector {
	panic("please build tags with postgres!")
}
