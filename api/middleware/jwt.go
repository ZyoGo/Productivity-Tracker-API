package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/w33h/formatter-response"
)

var (
	jwtSignedMethod = jwt.SigningMethodHS256
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("enter middleware")
			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, formatter.ForbiddenResponse("Invalid Token"))
			}

			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, formatter.ForbiddenResponse("Invalid Token"))
			}

			claim := jwt.MapClaims{}
			token, _ := jwt.ParseWithClaims(signature[1], claim, func(t *jwt.Token) (interface{}, error) {
				return []byte("Secret_JWT"), nil
			})

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || !token.Valid || method != jwtSignedMethod {
				return c.JSON(http.StatusForbidden, formatter.ForbiddenResponse("Invalid Token"))
			}

			return next(c)
		}
	}
}

func ExtractToken(c echo.Context) (id string, err error) {
	signature := strings.Split(c.Request().Header.Get("Authorization"), " ")

	claim := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(signature[1], claim, func(t *jwt.Token) (interface{}, error) {
		return []byte("Secret_JWT"), nil
	})

	//id = claim["UserId"]
	id = fmt.Sprintf("%v", claim["UserId"])

	return id, nil
}
