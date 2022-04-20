package controllers

import (
	"encoding/json"
	"govirt/configs"
	"govirt/models"
	"govirt/responses"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

//	@endpoint		POST /oauth/token
//	@desc				Authenticate with Ovirt server, and save token to database
func RequestToken(config *configs.Config) echo.HandlerFunc {
	return func (c echo.Context) error {
		client := config.HttpClient
		endpoint := configs.EnvOvirtUrl()	+ "/ovirt-engine/sso/oauth/token"
		data := url.Values{
			"grant_type": {"password"},
			"scope": 			{"ovirt-app-api"},
			"username": 	{c.FormValue("username")},
			"password": 	{c.FormValue("password")},
		}
	
		req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(data.Encode()))
		if err != nil {
			log.Printf("error creating HTTP request: %v", err)
			return c.JSON(http.StatusInternalServerError, responses.HTTPResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}
	
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Accept", "application/json")
	
		res, err := client.Do(req)
		if err != nil {
			log.Printf("error sending HTTP request: %v", err)
			return c.JSON(http.StatusInternalServerError, responses.HTTPResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}
		
		defer res.Body.Close()
	
		if res.StatusCode >= 200 && res.StatusCode <= 299 {
			tokenRes := responses.TokenResponse{}
	
			d := json.NewDecoder(res.Body)
			if err = d.Decode(&tokenRes); err != nil {
				log.Fatalf("error deserializing token data %v", err)
			}
	
			// Save username and token to database
			newUser := models.User{
				Ovirt: models.OvirtUser{Username: data.Get("username"), Access_token: tokenRes.AccessToken},
			}
	
			_, err := models.SaveUser(newUser, config.Db)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, responses.HTTPResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}
	
			return c.JSON(http.StatusOK, responses.HTTPResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": tokenRes}})
		} else {
			errorRes := responses.ErrorResponse{}
	
			d := json.NewDecoder(res.Body)
			if err = d.Decode(&errorRes); err != nil {
				log.Fatalf("error deserializing error data %v", err)
			}
	
			return c.JSON(res.StatusCode, responses.HTTPResponse{Status: res.StatusCode, Message: "error", Data: &echo.Map{"data": errorRes}})
		}
	}
}
