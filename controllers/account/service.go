package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "go_test/db"
	models "go_test/models"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func GetAccounts(c echo.Context) error {

	var data []models.Account
	db.DbManager().Model(models.Account{}).Find(&data)

	return c.JSON(http.StatusOK, data)
}

func GetAccount(c echo.Context) error {

	id := c.Param("id")

	var data []models.Account

	err := db.DbManager().Model(models.Account{}).Where("Id = ?", id).Find(&data)

	if err.Error != nil {
		c.JSON(http.StatusOK, map[string]interface{}{"status": true, "data": []string{}})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"status": true, "data": data[0]})
}

func CreateAccount(c echo.Context) error {

	fmt.Println(111)
	request_data := make(map[string]interface{})
	json.NewDecoder(c.Request().Body).Decode(&request_data)
	fmt.Println(222)

	id := uuid.New().String()
	fmt.Println(333, request_data)
	db.DbManager().Create(models.Account{
		Id:            id,
		UserId:        request_data["user_id"].(string),
		AccountNumber: request_data["account_number"].(string),
		Balance:       request_data["balance"].(float64),
	})
	fmt.Println(444)

	return c.JSON(http.StatusOK, map[string]interface{}{"status": true, "message": "Account for user created", "data": map[string]interface{}{"account_id": id}})
}

func UpdateAccount(c echo.Context) error {

	id := c.Param("id")
	request_data := make(map[string]interface{})
	json.NewDecoder(c.Request().Body).Decode(&request_data)

	db.DbManager().Model(models.Account{}).Where("Id = ?", id).Updates(request_data)

	return c.JSON(http.StatusOK, map[string]interface{}{"status": true, "message": "User Account details updated"})
}

func DeleteAccount(c echo.Context) error {
	id := c.Param("id")

	db.DbManager().Model(models.Account{}).Where("Id = ?", id).Delete(id)

	return c.JSON(http.StatusOK, map[string]interface{}{"status": true, "message": "User account deleted"})
}
