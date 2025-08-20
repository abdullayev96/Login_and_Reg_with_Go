package services

import (
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"

    "Login_Reg_Go/models"
    "Login_Reg_Go/repositories/user_repository"
)



func RegisterUser(db *gorm.DB, username string, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := models.User{Username: username, Password: string(hashedPassword)}
    return user_repository.CreateUser(db, &user)
}



func ValidateUserCredentials(db *gorm.DB, username string, password string) error {
    user, err := user_repository.GetUserByUsername(db, username)
    if err != nil {
        return err
    }

    return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}