package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-clean-arch-boilerplate/library"
	"github.com/labstack/echo/v4"
)

//Claims struct
type Claims struct {
	jwt.StandardClaims
}

//ClaimsData struct
type ClaimsData struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

var jwtKey = []byte("my_secret_key")

// ValidateToken func
func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var response library.HTTPResponse
		token := c.Request().Header.Get("Authorization")
		splitToken := strings.Split(token, "Bearer")
		token = splitToken[1]
		token = strings.TrimSpace(token)

		if token == "" {
			response.StatusCode = "400"
			response.Message = "token cant be null"
			response.Data = nil

			return c.JSON(http.StatusBadRequest, response)
		}

		_, verifyKey := library.InitKeys()

		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				response.StatusCode = "401"
				response.Message = "unauthorized"
				response.Data = nil

				return c.JSON(http.StatusUnauthorized, response)
			}

			//cek expired token
			v, _ := err.(*jwt.ValidationError)
			if v.Errors == jwt.ValidationErrorExpired {
				response.StatusCode = "401"
				response.Message = "token is expired"
				response.Data = nil

				return c.JSON(http.StatusUnauthorized, response)
			}

			response.StatusCode = "400"
			response.Message = "token not valid"
			response.Data = nil

			return c.JSON(http.StatusBadRequest, response)
		}

		if !tkn.Valid {
			response.StatusCode = "401"
			response.Message = "unauthorized"
			response.Data = nil

			return c.JSON(http.StatusUnauthorized, response)
		}

		claimsData := extractClaimsData(tkn)
		c.Response().Header().Set("userEmail", claimsData.Issuer)

		return next(c)
	}
}

func extractClaimsData(tkn *jwt.Token) ClaimsData {
	var claimsData ClaimsData

	byteClaims, err := json.Marshal(tkn.Claims)
	if err != nil {
		panic(err)
	}

	_ = json.Unmarshal(byteClaims, &claimsData)

	return claimsData
}
