package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/go-playground/validator/v10"

    "Login_Reg_Go/services"
)
var validate = validator.New()



//   Register

func Register(c *gin.Context, db *gorm.DB) {
    var temp struct {
        Username string `json:"username" validate:"required,min=3,max=20"`
        Password string `json:"password" validate:"required,min=8"`
    }

    if err := c.ShouldBindJSON(&temp); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }
    
    if err := validate.Struct(&temp); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
        return
    }

    if err := services.RegisterUser(db, temp.Username, temp.Password); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}


//    Login 


func Login(c *gin.Context, db *gorm.DB) {
    var temp struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&temp); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if err := services.ValidateUserCredentials(db, temp.Username, temp.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}