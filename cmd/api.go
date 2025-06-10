package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
	"go-fiber-starter/internal/api/auth"
	"go-fiber-starter/internal/middleware"
	"go-fiber-starter/pkg/config"
	"go-fiber-starter/pkg/logger"
	"os"
)

func api() {
	// 创建Fiber应用
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(fiberLogger.New(fiberLogger.Config{
		Format: "${time} ${ip} ${status} ${latency} ${method} ${path}\n",
		Output: os.Stdout,
	}))

	auth.RegisterUnProtectedRoutes(app)
	// 配置路由组
	api := app.Group("/api")
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Current.Jwt.Secret),
	}))

	auth.RegisterRoutes(api)

	if err := app.Listen(":" + config.Current.App.Port); err != nil {
		logger.Fatal("启动服务器失败: %v", err)
	} else {
		logger.Info("服务器启动成功: http://127.0.0.1:%v ", config.Current.App.Port)
	}
}
