package db

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "Login_Reg_Go/models"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
    database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    database.AutoMigrate(&models.User{})
    DB = database
    return database
}
