package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	model "go-fiber-starter/internal/model/User"
	"go-fiber-starter/pkg/config"
	"go-fiber-starter/pkg/db"
	"time"
)

func GenerateJWT(user *model.User) (string, error) {
	// 自定义声明：除了标准的 exp，还加载你的业务字段
	claims := jwt.MapClaims{
		"user_id":   user.Id,
		"user_name": user.Username,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Current.Jwt.Secret))
}

func CurrentUser(c *fiber.Ctx) (user *model.User, err error) {
	raw := c.Locals("user")
	if raw == nil {
		return nil, errors.New("no jwt token in context")
	}

	token := raw.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	userId, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("user_id claim missing")
	}

	dbUser, err := db.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return &dbUser, nil
}
