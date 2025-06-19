package middlewares

import (
	"errors"
	"lolymarsh/pkg/common"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (am *Middleware) IsHaveTokenMiddleware(allowedRoles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conf := am.conf

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return common.HandleError(c, errors.New("unauthorized: missing authorization header"), http.StatusUnauthorized)
			}

			if !strings.HasPrefix(authHeader, "Bearer ") {
				return common.HandleError(c, errors.New("unauthorized: invalid authorization header format"), http.StatusUnauthorized)
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			if tokenString == "" {
				return common.HandleError(c, errors.New("unauthorized: missing token"), http.StatusUnauthorized)
			}

			secretKey := []byte(conf.Auth.SecretKey)
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("invalid_signing_method")
				}
				return secretKey, nil
			})

			if err != nil || !token.Valid {
				return common.HandleError(c, errors.New("unauthorized: invalid or expired token"), http.StatusUnauthorized)
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return common.HandleError(c, errors.New("claims_error"), http.StatusBadRequest)
			}

			expirationTime, ok := claims["exp"].(float64)
			if !ok {
				return common.HandleError(c, errors.New("expiration_time_error"), http.StatusBadRequest)
			}

			expirationUnix := int64(expirationTime)
			expirationDate := time.Unix(expirationUnix, 0)

			if time.Now().After(expirationDate) {
				return common.HandleError(c, errors.New("token_has_expired"), http.StatusUnauthorized)
			}

			role, ok := claims["role"].(string)
			if !ok {
				return common.HandleError(c, errors.New("role_error"), http.StatusBadRequest)
			}
			roleAllowed := slices.Contains(allowedRoles, role)
			if !roleAllowed {
				return common.HandleError(c, errors.New("access_denied"), http.StatusForbidden)
			}

			c.Set("user_id", claims["user_id"])
			c.Set("first_name", claims["first_name"])
			c.Set("last_name", claims["last_name"])
			c.Set("email", claims["email"])
			c.Set("role", claims["role"])

			return next(c)
		}
	}
}
