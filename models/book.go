package models

//
//import (
//	"fmt"
//	"gorm.io/gorm"
//)
//
//type Books struct {
//	ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
//	Author    *string `json:"author"`
//	Title     *string `json:"title"`
//	Publisher *string `json:"publisher"`
//}
//
//func MigrateBooks(db *gorm.DB) (err error) {
//	err = db.AutoMigrate(&Books{})
//	return fmt.Errorf("failed with Automigrations: %s", err)
//}

//package models
//
//import "gorm.io/gorm"
//
//type Books struct {
//	ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
//	Author    *string `json:"author"`
//	Title     *string `json:"title"`
//	Publisher *string `json:"publisher"`
//}
//
//func MigrateBooks(db *gorm.DB) error {
//	err := db.AutoMigrate(&Books{})
//	return err
//}
