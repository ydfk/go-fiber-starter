/*
 * @Description: Copyright (c) ydfk. All rights reserved
 * @Author: ydfk
 * @Date: 2025-06-09 17:48:23
 * @LastEditors: ydfk
 * @LastEditTime: 2025-06-09 17:59:21
 */
package auth

import (
	"github.com/gofiber/fiber/v2"
	model "go-fiber-starter/internal/model/User"
	"go-fiber-starter/internal/service"
	"go-fiber-starter/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var req struct{ Username, Password string }

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "invalid payload"})
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := model.User{Username: req.Username, Password: string(hash)}
	if err := db.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "用户名已存在"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"msg": "registered"})
}

func Login(c *fiber.Ctx) error {
	var req struct{ Username, Password string }

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "invalid payload"})
	}

	var user model.User
	if err := db.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"msg": "invalid credentials"})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"msg": "invalid credentials"})
	}

	token, err := service.GenerateJWT(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": "token generation error"})
	}
	return c.JSON(fiber.Map{"token": token})
}

func Profile(c *fiber.Ctx) error {
	user, err := service.CurrentUser(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": "user not found"})
	}
	return c.JSON(user)
}
