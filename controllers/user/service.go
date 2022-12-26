package user

import (
	"encoding/json"
	"go_test/db"
	"go_test/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	var data []models.User
	db.DbManager().Preload("Accounts").Model(models.User{}).Find(&data)

	return c.JSON(http.StatusOK, data)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")

	var data []models.User

	err := db.DbManager().Where("Id = ?", id).Preload("Accounts").Model(models.User{}).Find(&data)

	if err.Error != nil {
		c.JSON(http.StatusOK, map[string]interface{}{"status": true, "data": []string{}})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"status": true, "data": data[0]})
}

func CreateUser(c echo.Context) error {
	request_data := make(map[string]interface{})
	json.NewDecoder(c.Request().Body).Decode(&request_data)

	id := uuid.New().String()
	db.DbManager().Create(models.User{
		Id:        id,
		UserName:  request_data["name"].(string),
		FirstName: request_data["first_name"].(string),
		LastName:  request_data["second_name"].(string),
	})
	return c.JSON(http.StatusOK, map[string]interface{}{"status": true, "message": "User created", "data": map[string]interface{}{"user_id": id}})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	request_data := make(map[string]interface{})
	json.NewDecoder(c.Request().Body).Decode(&request_data)

	db.DbManager().Model(models.User{}).Where("Id = ?", id).Updates(request_data)

	return c.JSON(http.StatusOK, map[string]interface{}{"status": true, "message": "User details updated"})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	db.DbManager().Model(models.User{}).Where("Id = ?", id).Delete(id)

	return c.JSON(http.StatusOK, map[string]interface{}{"status": true, "message": "User details deleted"})
}
