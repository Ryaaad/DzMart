package services

import (
	"DzMart/dtos"
	"DzMart/initializers"
	"DzMart/models"
)

func GetAllUsers() ([]models.User, error) {
	var Users []models.User
	result := initializers.DB.Find(&Users)
	return Users, result.Error
}

func GetUserById(id string) (*models.User, error) {
	var user *models.User
	result := initializers.DB.Preload("Fav").Preload("Comments").First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil

}

func CreateUser(body *models.User) error {
	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Credit:   0,
		Sub:      false,
		Password: body.Password,
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateUser(input dtos.UpdateUserInput, user *models.User) (*models.User, error) {
	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Password != nil {
		user.Password = *input.Password
	}

	SaveErr := initializers.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"Name":     user.Name,
		"Email":    user.Email,
		"Password": user.Password,
	})

	if SaveErr.Error != nil {
		return nil, SaveErr.Error
	}

	return user, nil
}

func DeleteUser(user *models.User) error {
	result := initializers.DB.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddFavorite(user *models.User, product *models.Product) (*models.User, error) {
	association := initializers.DB.Model(user).Association("Fav").Append(product)
	if association != nil {
		return nil, association
	}

	return user, nil
}

func DeleteFavorite(user models.User, product models.Product) error {
	result := initializers.DB.Model(&user).Association("Fav").Delete(product)
	if result != nil {
		return result
	}

	return nil
}

func DeleteAllFavorite(user models.User) error {
	result := initializers.DB.Model(&user).Association("Fav").Clear()
	if result != nil {
		return result
	}

	return nil
}
