/*
 * @Description: Copyright (c) ydfk. All rights reserved
 * @Author: ydfk
 * @Date: 2025-06-10 11:23:18
 * @LastEditors: ydfk
 * @LastEditTime: 2025-06-10 16:37:23
 */
package main

import (
	"go-fiber-starter/internal/api/auth"
	"go-fiber-starter/internal/middleware"
	"go-fiber-starter/pkg/config"
	"go-fiber-starter/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/swagger"
)

func api() {
	// 创建Fiber应用
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(fiberLogger.New(fiberLogger.Config{
		Format: "${ip} ${status} ${latency} ${method} ${path}\n",
		Output: logger.GetFiberLogWriter(),
	}))

	auth.RegisterUnProtectedRoutes(app)
	// 配置路由组
	api := app.Group("/api")
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Current.Jwt.Secret),
		// 添加自定义错误处理，返回401状态码
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			logger.Error("JWT验证失败: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"message": "认证失败，请先登录",
			})
		},
	}))

	auth.RegisterRoutes(api)

	if err := app.Listen(":" + config.Current.App.Port); err != nil {
		logger.Fatal("启动服务器失败: %v", err)
	} else {
		logger.Info("服务器启动成功: http://127.0.0.1:%v ", config.Current.App.Port)
	}
}
