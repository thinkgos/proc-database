//go:build !mysql
// +build !mysql

package database

import "gorm.io/gorm"

func NewMySQL(dsn string) gorm.Dialector {
	panic("please build tags with mysql!")
}
