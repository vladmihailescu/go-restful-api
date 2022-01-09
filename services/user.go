package services

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/vladmihailescu/go-restful-api/database"
	"github.com/vladmihailescu/go-restful-api/models"
	"github.com/vladmihailescu/go-restful-api/utils"
)

// RegisterUser saves the user in db and returns a valid jwt
func RegisterUser(userRegisterDTO models.UserRegisterDTO) (string, error) {
	db := database.DBConn

	userRegisterDTO.Password = utils.HashPassword(userRegisterDTO.Password)

	user := userRegisterDTO.ToModel()
	if err := db.Create(&user).Error; err != nil {
		return "", fmt.Errorf("unable to create user in db")
	}

	jwt, err := utils.GenerateToken(user.ID)
	if err != nil || user.ID == 0 {
		return "", fmt.Errorf("unable to create token")
	}

	return jwt, nil
}

// LoginUser validates the userLoginDTO agains the db and returns a jwt
func LoginUser(userLoginDTO models.UserLoginDTO) (string, error) {
	db := database.DBConn

	userLoginDTO.Password = utils.HashPassword(userLoginDTO.Password)

	var user models.User
	err := db.Where(models.User{
		Email:    userLoginDTO.Email,
		Password: userLoginDTO.Password,
	}).First(&user).Error
	if err != nil {
		return "", fmt.Errorf("unable to find user in db by username and password")
	}

	jwt, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("unable to create token")
	}

	return jwt, nil
}

// GetProfileUser returns the profile of the userId
func GetProfileUser(userId uint) (*models.UserDTO, error) {
	db := database.DBConn

	var user models.User
	if err := db.Where(gorm.Model{ID: userId}).Find(&user).Error; err != nil {
		return nil, fmt.Errorf("unable to find user in db by id")
	}

	return &models.UserDTO{Name: user.Name, Email: user.Email}, nil
}

// GetAllUsers returns a list of all users in db
func GetAllUsers() ([]models.UserDTO, error) {
	db := database.DBConn

	var ret []models.UserDTO

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("unable to find users in db")
	}

	for _, user := range users {
		userDTO := models.UserDTO{Name: user.Name, Email: user.Email}
		ret = append(ret, userDTO)
	}
	return ret, nil
}
