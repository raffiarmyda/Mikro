package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
	"time"
)

type JWTCustomClaims struct {
	ID      int  `json:"id"`
	IsAdmin bool `json:"is_admin"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		SuccessHandler: func(context echo.Context) {
			//userId := context.Get("userId")
		},
		//ErrorHandlerWithContext: func(e error, c echo.Context) error {
		//return deliveries.NewErrorResponse(c, http.StatusUnauthorized, businesses.ErrInvalidTokenCredential)
		//},
	}
}

func (jwtConf *ConfigJWT) GenerateTokenJWT(id int, isAdmin bool) (string, error) {
	claims := JWTCustomClaims{
		id,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, nil
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JWTCustomClaims)
		isAdmin := claims.IsAdmin
		if isAdmin == false {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func IsUserId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		userId := c.Param("userId")
		convUserId, err := strconv.Atoi(userId)
		if err != nil {
			return echo.ErrBadRequest
		}
		claims := user.Claims.(*JWTCustomClaims)
		claimUserId := claims.ID
		if claimUserId != convUserId {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
