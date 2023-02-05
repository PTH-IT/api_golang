package gormdb

import "gorm.io/gorm"

var DB *gorm.DB

func Start(gormdb *gorm.DB) {
	DB = gormdb
}
func (repo userRepository) Begin() *gorm.DB {
	return DB.Begin()
}
