package user_repository

import (
    "gorm.io/gorm"

    "codesignal.com/todoapp/models"
)



func CreateUser(db *gorm.DB, user *models.User) error {
    result := db.Create(user)
    return result.Error
}



func GetUserByUsername(db *gorm.DB, username string) (*models.User, error) {
    var user models.User
    err := db.Where("username = ?", username).First(&user).Error
    return &user, err
}