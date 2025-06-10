/*
 * @Description: Copyright (c) ydfk. All rights reserved
 * @Author: ydfk
 * @Date: 2025-06-09 17:48:23
 * @LastEditors: ydfk
 * @LastEditTime: 2025-06-10 16:52:35
 */
package auth

import (
	"go-fiber-starter/internal/api/response"
	model "go-fiber-starter/internal/model/User"
	"go-fiber-starter/internal/service"
	"go-fiber-starter/pkg/db"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var req struct{ Username, Password string }

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, "参数不正确")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := model.User{Username: req.Username, Password: string(hash)}
	if err := db.DB.Create(&user).Error; err != nil {
		return response.Error(c, "用户名已存在")
	}

	return response.Success(c, user)
}

// @Summary 用户登录
// @Description 用户登录接口
// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "登录信息"
// @Success 200 {object} LoginResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/auth/login [post]
func Login(c *fiber.Ctx) error {
	var req struct{ Username, Password string }

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, "参数不正确")
	}

	var user model.User
	if err := db.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return response.Error(c, "用户名不存在")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return response.Error(c, "密码不正确")
	}

	token, err := service.GenerateJWT(&user)
	if err != nil {
		return response.Error(c, "token生成失败")
	}
	return response.Success(c, fiber.Map{"token": token})
}

func Profile(c *fiber.Ctx) error {
	user, err := service.CurrentUser(c)
	if err != nil {
		return response.Error(c, "用户未找到")
	}
	return response.Success(c, user)
}
