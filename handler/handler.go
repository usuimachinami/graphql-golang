package handler

import (
	"bytes"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"graphql-golang/common"
	"graphql-golang/db"
	"graphql-golang/gql"
	"graphql-golang/model"
	"log"
	"net/http"
	"time"
	"strconv"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		db := db.ConnectGORM()
		user := []model.User{}
		db.Where("name=? and password=?", username, password).Find(&user)

		if len(user) > 0 && username == user[0].Name {
			// Create token
			token := jwt.New(jwt.SigningMethodHS256)

			// Set claims
			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = username
			claims["admin"] = true
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			// Generate encoded token and send it as response.
			t, err := token.SignedString([]byte(common.SECRET_KEY))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{
				"user_id": strconv.FormatInt(user[0].Id, 10),
				"token": t,
			})
		}

		return echo.ErrUnauthorized
	}
}

func Query() echo.HandlerFunc {
	return func(c echo.Context) error {
		bufBody := new(bytes.Buffer)
		bufBody.ReadFrom(c.Request().Body)
		query := bufBody.String()
		log.Printf(query)
		result := gql.ExecuteQuery(query)
		return c.JSON(http.StatusOK, result)
	}
}
